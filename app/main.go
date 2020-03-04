package main

import (
	"app/modules/scraper"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	scraper.RegisterSpiderServiceServer(server, &scraper.Scraper{})

	log.Printf("Service Running on: 0.0.0.0:50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Server fail %v", err)
	}
}
