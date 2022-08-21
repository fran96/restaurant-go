package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/fran96/restaurant-go/contracts"
	waiter "github.com/fran96/restaurant-go/internal/waiter"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWaiterServiceServer(s, &waiter.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// func main() {
// 	cfg, err := config.ParseConfig()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	s := server.NewServer(cfg)
// 	s.Run()
// }
