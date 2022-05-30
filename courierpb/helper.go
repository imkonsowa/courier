package courierpb

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"courier/pkg/env"
)

func NewCourierClient() (CourierServiceClient, *grpc.ClientConn) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	c, err := grpc.Dial(env.String("GRPC_HOST", "localhost:1997"), opts)
	if err != nil {
		log.Panicf("could not connect: %v", err)
	}

	client := NewCourierServiceClient(c)

	return client, c
}
