package api

import (
	"fmt"
	"gmaps-service/config"
	"gmaps-service/proto"
	"log"
	"time"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

type ServerGRPC struct {
	config config.Config
	proto.UnimplementedLocationServiceServer
}

func NewGrpcServer(config config.Config) ServerGRPC {
	return ServerGRPC{
		config: config,
	}
}

func (s *ServerGRPC) FindClosest(ctx context.Context, in *proto.LocationRequest) (*proto.Location, error) {

	// No destinations means we can't return anything.
	if len(in.Destinations) == 0 {
		return nil, nil
	}

	// Intialize Google Maps API.
	client, err := maps.NewClient(maps.WithAPIKey(s.config.MapsKey))
	if err != nil {
		log.Println("Unable to connect to Google Maps API: ", err)
		return nil, err
	}

	var destinations []string
	for _, loc := range in.Destinations {
		destinations = append(destinations, fmt.Sprintf("%f,%f", loc.Lat, loc.Lng)) // note the = instead of :=
	}

	req := &maps.DistanceMatrixRequest{
		Origins:      []string{fmt.Sprintf("%f,%f", in.Origin.Lat, in.Origin.Lng)},
		Destinations: destinations,
		Units:        "UnitsMetric",
	}

	// Call Google Maps external API.
	matrix, err := client.DistanceMatrix(context.Background(), req)
	if err != nil {
		log.Println("Error retrieving data from Google Maps API: ", err)
		return nil, err
	}

	var minIndex int = 0
	var minDuration time.Duration = 1000 * time.Hour

	// Find closest station.
	for i, elem := range matrix.Rows[0].Elements {
		if elem.Duration < minDuration {
			minIndex = i
			minDuration = elem.Duration
		}
	}

	return in.Destinations[minIndex], nil
}
