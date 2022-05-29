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
)

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

	// ReadAll instead of Reading line by line eliminates the complexity
	// As one file is received daily to the service, it will not a memory critical issue.
	all, err := csv.NewReader(file).ReadAll()
	if len(all) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Empty files not allowed",
		})
		return
	}

	if len(all) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Empty files not allowed",
		})
		return
	}

	fmt.Println("Start processing file:", header.Filename)

	go h.processLines(all)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Well received, Processing!",
	})
	return
}

func (h *CsvHandler) processLines(lines [][]string) {
	stream, streamErr := h.client.ProcessParcels(context.Background())
	if streamErr != nil {
		return
	}

	start := 1
	end := 1000

	for {
		if end >= len(lines) {
			end = len(lines) - 1
		}

		payload := lines[start:end]

		var parcels []*courierpb.Parcel
		for _, parcel := range payload {
			id, err := strconv.Atoi(removeSpaces(parcel[0]))
			if err != nil {
				continue
			}
			weight, weightErr := strconv.ParseFloat(removeSpaces(parcel[3]), 8)
			if weightErr != nil {
				continue
			}

			parcels = append(parcels, &courierpb.Parcel{
				Id:     int64(id),
				Email:  parcel[1],
				Phone:  parcel[2],
				Weight: float32(weight),
			})
		}

		request := &courierpb.ProcessParcelsRequest{
			Date:    time.Now().Format("2006-01-02"),
			Parcels: parcels,
		}

		err := stream.Send(request)
		if err != nil {
			return
		}

		if end >= len(lines)-1 {
			break
		}

		end += 1000
		start += 1000
	}

	waitc := make(chan struct{})
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}
			fmt.Printf("Received: %v\n", res.GetMessage())
		}
		close(waitc)
	}()

	<-waitc

	stream.CloseSend()
}

func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
