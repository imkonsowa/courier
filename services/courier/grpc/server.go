package grpc

import (
	"fmt"
	"io"
	"log"

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
			log.Fatalf("Error processing parcels stream from client, err: %v", err)
			return err
		}

		var upsert []*models.Parcel
		for _, parcel := range req.GetParcels() {
			upsert = append(upsert, models.FromPb(parcel))
		}
		insert, err := s.ParcelAdapter.PatchInsert(upsert)
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Inserted rows count: %d", len(insert)))

		replyErr := stream.Send(&courierpb.ProcessParcelsResponse{
			Message: fmt.Sprintf("Received and processing: %d", req.GetParcels()[len(req.GetParcels())-1].GetId()),
		})
		if replyErr != nil {
			log.Fatalf("Failed to reply to the recieved parcels, err %v", replyErr)
			return replyErr
		}
	}
}
