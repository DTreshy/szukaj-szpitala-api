package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DTreshy/szukaj-szpitala-api/api/config"
	"github.com/DTreshy/szukaj-szpitala-api/api/database"
	"github.com/DTreshy/szukaj-szpitala-api/api/grpc"
	googleRPC "google.golang.org/grpc"
)

func main() {
	cfg, db, err := setUpService()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Client.Disconnect(context.Background())

	grpcClose := make(chan struct{})
	grpcServer := grpc.NewGrpcServer(db)

	go func() {
		if err := runGrpcServer(grpcServer, cfg.GrpcBind); err != nil {
			log.Println(err)
		}

		close(grpcClose)
	}()

	defer grpcServer.GracefulStop()

	interruptC := make(chan os.Signal, 1)

	signal.Notify(interruptC, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		signal.Stop(interruptC)
		close(interruptC)
	}()

	select {
	case <-grpcClose:
	case <-interruptC:
	}
}

func setUpService() (*config.Config, *database.Db, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		return nil, nil, err
	}

	db, err := database.NewDb(cfg.Username, cfg.Password, cfg.MongoUri)
	if err != nil {
		return nil, nil, err
	}

	if err := db.Connect(ctx); err != nil {
		return nil, nil, err
	}

	return cfg, db, nil
}

func runGrpcServer(s *googleRPC.Server, bind string) error {
	lis, err := net.Listen("tcp", bind)
	if err != nil {
		return err
	}

	return s.Serve(lis)
}
