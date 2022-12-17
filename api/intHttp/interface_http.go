package intHttp

import (
	"context"
	"net/http"
	"time"

	"github.com/DTreshy/szukaj-szpitala-api/api/intGrpc/proto"
	lmiddleware "github.com/DTreshy/szukaj-szpitala-api/api/intHttp/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewHttp(ctx context.Context, debug bool, bind, grpcAddress string, trustedOrigins []string) (*http.Server, error) {
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	gmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := proto.RegisterClientHandlerFromEndpoint(ctx, gmux, grpcAddress, opts)
	if err != nil {
		return nil, err
	}

	handler := lmiddleware.CORS(trustedOrigins)

	return &http.Server{
		Addr:              bind,
		Handler:           handler(gmux),
		ReadHeaderTimeout: 15 * time.Second,
	}, nil
}
