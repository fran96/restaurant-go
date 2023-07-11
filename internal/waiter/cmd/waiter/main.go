package main

import (
	"flag"
	"log"
	"net"

	waiterPb "github.com/fran96/restaurant-go/contracts"
	"github.com/fran96/restaurant-go/internal/util"
	"github.com/fran96/restaurant-go/internal/waiter"
	"google.golang.org/grpc"
)

func main() {
	go func() {
		// consume
		waiter.Consume()
	}()

	flag.Parse()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	lis, err := net.Listen("tcp", config.WaiterServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	waiterPb.RegisterWaiterServiceServer(s, &waiter.Server{})
	log.Printf("waiter server listening at %v", config.WaiterServerAddress)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	s.GracefulStop()
	log.Print("Server Exited Properly")

}
