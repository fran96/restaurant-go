package kitchen

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/fran96/restaurant-go/contracts/kitchen"
)

// Server represents the kitchen gRPC server
type Server struct {
	pb.UnimplementedKitchenServiceServer
}

func (ks Server) Cook(ctx context.Context, in *pb.MakeFood) (*pb.OrderReceived, error) {
	if in.ListOfFood == nil {
		return nil, errors.New("can't cook anything")
	}

	fmt.Printf("Kitchen - OrderReceived: %v ", in.Id)
	return &pb.OrderReceived{Message: fmt.Sprintf("Order received ID: %s", in.Id)}, nil
}
