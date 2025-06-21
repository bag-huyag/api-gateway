package client

import (
	"log"

	pb "github.com/bag-huyag/api-gateway/proto/gen"
	"google.golang.org/grpc"
)

func NewUserServiceClient(addr string) pb.UserServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user-service: %v", err)
	}
	return pb.NewUserServiceClient(conn)
}
