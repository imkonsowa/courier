package handlers

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"courier/courierpb"
	"courier/services/csv_parser/utils"
)

const ChunkToProcessThreshold = 1000

type CsvHandler struct {
	client courierpb.CourierServiceClient
}

func NewCsvHandler(c courierpb.CourierServiceClient) *CsvHandler {
	return &CsvHandler{
		client: c,
	}
}

func (h *CsvHandler) ProcessParcels(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to process your request",
		})
		return
	}

	if file == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "File is missing!",
		})
		return
	}

	all, err := csv.NewReader(file).ReadAll()
	if len(all) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Empty files not allowed",
		})
		return
	}

	fmt.Println("Start processing file:", header.Filename)

	go func() {
		err := h.processLines(all)
		if err != nil {
			log.Printf(err.Error())
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Well received, Processing!",
	})
	return
}

func (h *CsvHandler) processLines(lines [][]string) error {
	// drop the header
	lines = lines[1:]

	stream, streamErr := h.client.ProcessParcels(context.Background())
	if streamErr != nil {
		return fmt.Errorf("failed to open a stream connection with courier service, err: %v", streamErr)
	}

	// a flag to indicate how many routine workers coordinating in sending data over the stream
	var chunksCount = 0
	chunksCount = len(lines) / ChunkToProcessThreshold
	if len(lines)-(chunksCount*ChunkToProcessThreshold) > 0 {
		chunksCount++
	}

	// worker coordinators channels
	lch := make(chan [][]string, chunksCount)
	sendch := make(chan interface{}, chunksCount)

	// ignite 10 workers to send chunks over the stream
	for i := 0; i < 10; i++ {
		go streamWorker(stream, lch, sendch)
	}

	for i := 1; i <= chunksCount; i++ {
		end := i * ChunkToProcessThreshold
		start := end - ChunkToProcessThreshold

		// safe check to mitigate out of range exception
		if end >= len(lines) {
			end = len(lines)
		}

		// reinitialize payload to avoid cap increase
		var payload [][]string
		payload = append(payload, lines[start:end]...)
		lch <- payload
	}

	// no more payloads to process
	close(lch)

	// a channel to block till all
	waitc := make(chan struct{})

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetMessage())
		}

		close(waitc)
	}()

	// assert that all chunks were sent
	// TODO: add error reporting from workers?
	for j := 0; j < chunksCount; j++ {
		<-sendch
	}
	stream.CloseSend()

	<-waitc

	fmt.Println("Wait is over")

	return nil
}

func streamWorker(
	stream courierpb.CourierService_ProcessParcelsClient,
	linesChannel <-chan [][]string,
	sendch chan<- interface{},
) {
	today := time.Now().Format("2006-01-02")

	for payload := range linesChannel {
		var parcels []*courierpb.Parcel

		for _, parcel := range payload {
			id, err := strconv.Atoi(utils.RemoveSpaces(parcel[0]))
			if err != nil {
				continue
			}
			weight, weightErr := strconv.ParseFloat(utils.RemoveSpaces(parcel[3]), 32)
			if weightErr != nil {
				continue
			}

			parcels = append(parcels, &courierpb.Parcel{
				Id:      int64(id),
				Email:   utils.RemoveSpaces(parcel[1]),
				Phone:   strings.TrimSpace(parcel[2]),
				Weight:  float32(weight),
				Country: utils.CountryFromPhone(parcel[2]).String(),
			})
		}

		request := &courierpb.ProcessParcelsRequest{
			Date:    today,
			Parcels: parcels,
		}

		err := stream.Send(request)
		if err != nil {
			log.Printf("failed send request to courier service, err: %v", err)
		}

		sendch <- nil
	}
}
