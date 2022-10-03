package waiter

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"

	pb "github.com/fran96/restaurant-go/contracts"
	pbKitchen "github.com/fran96/restaurant-go/contracts/kitchen"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

// Server represents the waiter gRPC server
type Server struct {
	pb.UnimplementedWaiterServiceServer
}

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
)

func (ws *Server) Order(ctx context.Context, in *pb.OrderRequest) (*pb.OrderAcknowledged, error) {

	id := uuid.New()
	fmt.Println("orderID: ", id.String())

	if len(in.ListOfFood) > 0 {
		fmt.Println("order rpc - contains food")

		// Call Kitchen
		makeFood := &pbKitchen.MakeFood{
			Id:         id.String(),
			ListOfFood: in.ListOfFood,
		}

		// Set up a connection to the server.
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pbKitchen.NewKitchenServiceClient(conn)
		orderReceived, err := client.Cook(ctx, makeFood)
		if err != nil {
			return nil, err
		}

		fmt.Println("OrderReceived from kitchen: ", orderReceived)

	}

	return &pb.OrderAcknowledged{OrderID: id.String()}, nil
}
