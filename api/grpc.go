package api

import (
	"gmaps-service/proto"
	"log"

	"golang.org/x/net/context"
)

type ServerGRPC struct {
	proto.UnimplementedLocationServiceServer
}

func (s *ServerGRPC) FindClosest(ctx context.Context, in *proto.LocationRequest) (*proto.Location, error) {
	log.Printf("ID = %d, lat = %.3f, lng = %.3f", in.MyLocation.Id, in.MyLocation.Lat, in.MyLocation.Lng)
	return &proto.Location{
		Id:  42,
		Lat: 4.20,
		Lng: 4.20,
	}, nil
}
