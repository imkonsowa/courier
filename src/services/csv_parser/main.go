package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"courier/src/courierpb"
	"courier/src/services/csv_parser/handlers"
)

func main() {
	// construct new gin engine with recovery and logger middleware attached
	engine := gin.Default()

	// disable tls for the gRPC server
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	c, err := grpc.Dial("localhost:1997", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	client := courierpb.NewCourierServiceClient(c)

	csvHandler := handlers.NewCsvHandler(client)

	engine.POST("/upload", csvHandler.ProcessParcels)

	shutdown := make(chan os.Signal)

	srv := &http.Server{
		Addr:    ":1996",
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to run a server. err: %v", err)
		}
	}()

	signal.Notify(shutdown, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)

	<-shutdown

	c.Close()
	srv.Close()
}
