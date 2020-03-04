package main

import (
	"app/modules/scraper"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello world")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	scraper.RegisterSpiderServiceServer(server, &scraper.Scraper{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Server fail %v", err)
	}
}
