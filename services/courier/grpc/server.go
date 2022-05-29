package grpc

import (
	"fmt"
	"io"
	"log"
	"time"

	"courier/courierpb"
	"courier/services/courier/data/adapters"
	"courier/services/courier/data/models"
)

type Server struct {
	ParcelAdapter *adapters.ParcelAdapter
}

func NewServer(pa *adapters.ParcelAdapter) *Server {
	return &Server{
		ParcelAdapter: pa,
	}
}

func (s *Server) ProcessParcels(stream courierpb.CourierService_ProcessParcelsServer) error {
	fmt.Println("ProcessParcels rpc received")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error processing parcels stream from client, err: %v", err)
			return err
		}

		date, _ := time.Parse("2006-01-02", req.GetDate())
		var upsert []*models.Parcel
		for _, parcel := range req.GetParcels() {
			upsert = append(upsert, models.FromPb(parcel, &date))
		}
		_, err = s.ParcelAdapter.PatchInsert(upsert)
		if err != nil {
			return err
		}

		replyErr := stream.Send(&courierpb.ProcessParcelsResponse{
			Message: fmt.Sprintf("Received and processing: %d", req.GetParcels()[len(req.GetParcels())-1].GetId()),
		})
		if replyErr != nil {
			log.Printf("Failed to reply to the recieved parcels, err %v", replyErr)
			return replyErr
		}
	}
}
