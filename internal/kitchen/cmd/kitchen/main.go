package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/fran96/restaurant-go/contracts/kitchen"
	kitchen "github.com/fran96/restaurant-go/internal/kitchen"
	"github.com/fran96/restaurant-go/internal/util"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	lis, err := net.Listen("tcp", config.KitchenServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterKitchenServiceServer(s, &kitchen.Server{})
	log.Printf("kitchen server listening at %v", config.KitchenServerAddress)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
