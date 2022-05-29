package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"courier/courierpb"
	"courier/pkg/config"
	"courier/services/courier/data/adapters"
	"courier/services/courier/data/models"
	grpcpkg "courier/services/courier/grpc"
)

func main() {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=GMT",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port,
		cfg.DB.SSlMode,
	)
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(fmt.Sprintf("failed to connect to DB; %s", dbErr))
	}
	migrateErr := db.AutoMigrate(&models.Parcel{})
	if migrateErr != nil {
		panic(fmt.Sprintf("failed to migrate; %s", migrateErr))
	}

	lis, err := net.Listen("tcp", "0.0.0.0:1997")
	if err != nil {
		log.Fatalf("TCP failed to listen, err: %v", err)
	}

	server := grpc.NewServer(
		grpc.StreamInterceptor(
			grpcmiddleware.ChainStreamServer(
				grpcrecovery.StreamServerInterceptor(),
			),
		),
	)

	s := grpcpkg.NewServer(
		adapters.NewParcelAdapter(db),
	)
	courierpb.RegisterCourierServiceServer(server, s)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Failed to start grpc server, err: %v", err)
		}
	}()

	fmt.Println("Listening and serving gRPC on :1997")

	shutdown := make(chan os.Signal)
	signal.Notify(
		shutdown,
		syscall.SIGINT,
	)
	<-shutdown

	server.GracefulStop()
	lis.Close()
}
