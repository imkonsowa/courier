package main

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"

	"courier/src/courierpb"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:1997")
	if err != nil {
		log.Fatalf("TCP failed to listen, err: %v", err)
	}

	server := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_recovery.StreamServerInterceptor(),
			),
		),
	)

	courierpb.RegisterCourierServiceServer(server, &Server{})

	watchCat := make(chan string)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Failed to start grpc server, err: %v", err)
		}
	}()

	fmt.Println("Listening and serving gRPC on :1997")

	<-watchCat
}
