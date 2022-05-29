package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"courier/courierpb"
	"courier/pkg/env"
	"courier/services/csv_parser/handlers"
)

func main() {
	engine := gin.Default()

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	c, err := grpc.Dial(env.String("GRPC_HOST", "localhost:1997"), opts)
	if err != nil {
		log.Panicf("could not connect: %v", err)
	}

	client := courierpb.NewCourierServiceClient(c)

	csvHandler := handlers.NewCsvHandler(client)

	engine.POST("/upload", csvHandler.ProcessParcels)

	shutdown := make(chan os.Signal)

	if err := engine.Run(":1996"); err != nil {
		log.Printf("Failed to run a server. err: %v", err)
	}

	signal.Notify(shutdown, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)

	<-shutdown

	c.Close()
}
