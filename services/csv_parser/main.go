package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"courier/courierpb"
	"courier/services/csv_parser/handlers"
)

func main() {
	engine := gin.Default()

	client, connection := courierpb.NewCourierClient()

	csvHandler := handlers.NewCsvHandler(client)

	engine.POST("/upload", csvHandler.ProcessParcels)

	shutdown := make(chan os.Signal)

	if err := engine.Run(":1996"); err != nil {
		log.Printf("Failed to run a server. err: %v", err)
	}

	signal.Notify(shutdown, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)

	<-shutdown

	connection.Close()
}
