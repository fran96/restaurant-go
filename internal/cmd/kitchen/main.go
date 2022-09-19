package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/fran96/restaurant-go/contracts/kitchen"
	kitchen "github.com/fran96/restaurant-go/internal/kitchen"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "The kitchen server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterKitchenServiceServer(s, &kitchen.Server{})
	log.Printf("kitchen server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
