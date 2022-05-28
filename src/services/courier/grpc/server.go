package grpc

import (
	"fmt"
	"io"
	"log"

	"courier/src/courierpb"
	"courier/src/services/courier/data/adapters"
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

		replyErr := stream.Send(&courierpb.ProcessParcelsResponse{
			Message: fmt.Sprintf("Received and processing: %d", req.GetParcels()[len(req.GetParcels())-1].GetId()),
		})
		if replyErr != nil {
			log.Fatalf("Failed to reply to the recieved parcels, err %v", replyErr)
			return replyErr
		}
	}
}
