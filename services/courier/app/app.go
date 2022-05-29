package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"courier/courierpb"
	"courier/services/courier/config"
	"courier/services/courier/data/adapters"
	"courier/services/courier/data/models"
	grpcpkg "courier/services/courier/grpc"
)

type App struct {
	Config *config.Config
	Engine *gin.Engine
	DB     *gorm.DB
}

func NewApp() *App {
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

	engine := gin.Default()

	return &App{
		DB:     db,
		Engine: engine,
		Config: cfg,
	}
}

func (a *App) Run() {
	if a.Engine == nil {
		panic("server engine is not constructed yet")
	}

	lis, err := net.Listen("tcp", "0.0.0.0:1997")
	if err != nil {
		log.Printf("TCP failed to listen, err: %v", err)
		return
	}

	server := grpc.NewServer(
		grpc.StreamInterceptor(
			grpcmiddleware.ChainStreamServer(
				grpcrecovery.StreamServerInterceptor(),
			),
		),
	)
	// TODO: refactor to an inversion provider
	s := grpcpkg.NewServer(
		adapters.NewParcelAdapter(a.DB),
	)
	courierpb.RegisterCourierServiceServer(server, s)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Printf("Failed to start the grpc server, err: %v", err)
		}

		fmt.Println("Listening and serving gRPC on :1997")
	}()

	go func() {
		if err := a.Engine.Run(":1998"); err != nil {
			log.Printf("Failed to start the http server, err: %v", err)
		}
	}()

	shutdown := make(chan os.Signal)
	signal.Notify(
		shutdown,
		syscall.SIGINT,
	)
	<-shutdown

	server.GracefulStop()
	lis.Close()
}
