package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DTreshy/szukaj-szpitala-api/api/intGrpc/proto"
	googleRPC "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type connection struct {
	conn           *googleRPC.ClientConn
	commandTimeout time.Duration
	rpc            proto.ClientClient
}

func newConnection() *connection {
	return &connection{
		conn:           nil,
		commandTimeout: 4 * time.Second,
		rpc:            nil,
	}
}

func (c *connection) connect(endpoint string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.commandTimeout)

	defer cancel()

	c.conn, err = googleRPC.DialContext(ctx, endpoint, googleRPC.WithTransportCredentials(insecure.NewCredentials()), googleRPC.WithBlock())
	if err != nil {
		return fmt.Errorf("cannot connect to %s: %s", endpoint, err)
	}

	c.rpc = proto.NewClientClient(c.conn)

	return nil
}

func (c *connection) sendInsertRequest(ctx context.Context, rec *record) error {
	newCtx, cancel := context.WithTimeout(ctx, c.commandTimeout)

	defer cancel()

	response, err := c.rpc.InsertHospital(newCtx, &proto.InsertHospitalRequest{
		Name:    rec.hospital.Name,
		City:    rec.hospital.City,
		Address: rec.hospital.Address,
		Phone:   rec.hospital.Phone,
		Email:   rec.hospital.Email,
	})
	if err != nil {
		return err
	}

	if response.Message != "successfully inserted hospital" {
		return fmt.Errorf("%s", response.Message)
	}

	return nil
}
