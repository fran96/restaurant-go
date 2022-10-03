package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/fran96/restaurant-go/contracts"
	"github.com/fran96/restaurant-go/internal/waiter"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	go func() {
		// consume
		waiter.Consume()
	}()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterWaiterServiceServer(s, &waiter.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
