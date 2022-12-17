package grpc

import (
	"context"
	"fmt"

	"github.com/DTreshy/szukaj-szpitala-api/api/database"
	"github.com/DTreshy/szukaj-szpitala-api/api/grpc/proto"
	"github.com/DTreshy/szukaj-szpitala-api/api/serialization"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	googleRPC "google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedClientServer
	db *database.Db
}

func NewGrpcServer(db *database.Db) *googleRPC.Server {
	s := googleRPC.NewServer(
		googleRPC.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		googleRPC.ChainUnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	proto.RegisterClientServer(s, &server{
		db: db,
	})
	grpc_prometheus.Register(s)

	return s
}

func (s *server) InsertHospital(ctx context.Context, r *proto.InsertHospitalRequest) (*proto.InsertHospitalReply, error) {
	if r.Address == "" {
		return &proto.InsertHospitalReply{Message: "address cannot be blank"}, nil
	}

	if r.City == "" {
		return &proto.InsertHospitalReply{Message: "city cannot be blank"}, nil
	}

	if r.Email == "" {
		return &proto.InsertHospitalReply{Message: "email cannot be blank"}, nil
	}

	if r.Name == "" {
		return &proto.InsertHospitalReply{Message: "name cannot be blank"}, nil
	}

	if r.Phone == "" {
		return &proto.InsertHospitalReply{Message: "phone cannot be blank"}, nil
	}

	hospital, err := serialization.NewHospital(r.Name, r.Address, r.City, r.Phone, r.Email)
	if err != nil {
		return &proto.InsertHospitalReply{Message: fmt.Sprintf("cannot serialize hospital: %s", err.Error())}, nil
	}

	if err := s.db.InsertHospital(ctx, hospital); err != nil {
		return &proto.InsertHospitalReply{Message: fmt.Sprintf("cannot insert hospital %s", err.Error())}, nil
	}

	return &proto.InsertHospitalReply{Message: "successfully inserted hospital"}, nil
}

func (s *server) QueryNearestHospitals(ctx context.Context, r *proto.QueryNearestHospitalsRequest) (*proto.QueryNearestHospitalsReply, error) {
	place := serialization.NewPlaceFromCoords(float64(r.Latitude), float64(r.Longitude))

	hospitals, err := s.db.QueryNearestHospitals(ctx, place, int(r.NumberHospitals))
	if err != nil {
		return &proto.QueryNearestHospitalsReply{Message: fmt.Sprintf("error quering hospitals: %s", err.Error())}, nil
	}

	protoHospitals := []*proto.Hospital{}

	for _, hospital := range hospitals {
		protoHospitals = append(protoHospitals, &proto.Hospital{
			Name:     hospital.Hospital.Name,
			Email:    hospital.Hospital.Email,
			Phone:    hospital.Hospital.Phone,
			City:     hospital.Hospital.Place.City,
			Address:  hospital.Hospital.Place.Address,
			Distance: float32(hospital.Distance),
		})
	}

	return &proto.QueryNearestHospitalsReply{
		NearestHospitals: protoHospitals,
		Message:          "Successfully queried hospitals",
	}, nil
}
