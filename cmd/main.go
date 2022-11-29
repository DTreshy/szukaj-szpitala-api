package main

import (
	"net"

	"github.com/DTreshy/szukaj-szpital-api/api/grpc"
	googleRPC "google.golang.org/grpc"
)

func main() {
	grpcClose := make(chan struct{})
	grpcServer := grpc.NewGrpcServer()

	go func() {
		runGrpcServer(grpcServer, "localhost:55011")
		close(grpcClose)
	}()

	defer grpcServer.GracefulStop()
}

func runGrpcServer(s *googleRPC.Server, bind string) error {
	lis, err := net.Listen("tcp", bind)
	if err != nil {
		return err
	}

	return s.Serve(lis)
}
