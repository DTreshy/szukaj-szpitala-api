package grpc

import (
	"github.com/DTreshy/szukaj-szpital-api/api/grpc/proto"
	googleRPC "google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedClientServer
}

func NewGrpcServer() *googleRPC.Server {
	s := googleRPC.NewServer()

	return s
}
