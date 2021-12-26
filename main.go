package main

import (
	"gmaps-service/api"
	"gmaps-service/config"
	"gmaps-service/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// Load configuration settings.
	config, err := config.New(".")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Create a server and setup routes.
	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("Failed to create a server: ", err)
	}

	// Run gRPC server concurrently.
	go func() {
		lis, err := net.Listen("tcp", config.GrpcAddress)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		s := api.ServerGRPC{}
		grpcServer := grpc.NewServer()

		proto.RegisterLocationServiceServer(grpcServer, &s)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	// Start a server.
	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("Failed to start a server: ", err)
	}
}
