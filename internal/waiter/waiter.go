package waiter

import (
	"flag"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"

	pb "github.com/fran96/restaurant-go/contracts"
	pbKitchen "github.com/fran96/restaurant-go/contracts/kitchen"
	"github.com/fran96/restaurant-go/internal/util"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

// Server represents the waiter gRPC server
type Server struct {
	pb.UnimplementedWaiterServiceServer
}

func (ws *Server) Order(ctx context.Context, in *pb.OrderRequest) (*pb.OrderAcknowledged, error) {

	id := uuid.New()

	if len(in.ListOfFood) > 0 {
		flag.Parse()
		config, err := util.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config:", err)
		}

		// Set up a connection to the server.
		conn, err := grpc.Dial(config.KitchenServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pbKitchen.NewKitchenServiceClient(conn)
		makeFood := &pbKitchen.MakeFood{
			Id:         id.String(),
			ListOfFood: in.ListOfFood,
		}

		orderReceived, err := client.Cook(ctx, makeFood)
		if err != nil {
			return nil, err
		}

		fmt.Printf("\nOrderReceived from kitchen %v at %v: ", orderReceived, time.Now())

	}

	return &pb.OrderAcknowledged{OrderID: id.String()}, nil
}
