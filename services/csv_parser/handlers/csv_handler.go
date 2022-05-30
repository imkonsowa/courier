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
	"courier/pkg/responses"
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
		responses.NewContextResponse(c).
			Error().
			Code(http.StatusInternalServerError).
			Message("Failed to process your request").
			Send()
		return
	}

	if file == nil {
		responses.NewContextResponse(c).
			Error().
			Code(http.StatusBadRequest).
			Message("File is missing!").
			Send()
		return
	}

	all, err := csv.NewReader(file).ReadAll()
	if len(all) == 0 {
		responses.NewContextResponse(c).
			Error().
			Code(http.StatusBadRequest).
			Message("Empty files not allowed").
			Send()
		return
	}

	fmt.Println("Start processing file:", header.Filename)

	go func() {
		err := h.processLines(all)
		if err != nil {
			log.Printf(err.Error())
		}
	}()

	responses.NewContextResponse(c).
		Success().
		Message("Well received, Processing!").
		Send()
}

func (h *CsvHandler) processLines(lines [][]string) error {
	// drop the header
	lines = lines[1:]

	stream, streamErr := h.client.ProcessParcels(context.Background())
	if streamErr != nil {
		return fmt.Errorf("failed to open a stream connection with courier service, err: %v", streamErr)
	}

	// a flag to indicate how many chunks expected to be sent over the stream.
	var chunksCount = 0
	chunksCount = len(lines) / ChunkToProcessThreshold
	if chunksCount < 0 {
		chunksCount = 1
	} else if len(lines)-(chunksCount*ChunkToProcessThreshold) > 0 {
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

	// a channel to block till all parcels sent
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

	// TODO: add error reporting from workers?
	for j := 1; j <= chunksCount; j++ {
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
