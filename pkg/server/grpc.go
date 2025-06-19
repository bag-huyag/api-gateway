package server

import (
	"fmt"
	"net"

	"github.com/bag-huyag/api-gateway/internal/handler"
	pb "github.com/bag-huyag/api-gateway/proto/gen"
	"google.golang.org/grpc"
)

func StartGRPC() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	grpcServer := grpc.NewServer()
	if grpcServer == nil {
		return fmt.Errorf("failed to create grpc server: %w", err)
	}
	pb.RegisterUserServiceServer(grpcServer, &handler.UserHandler{})

	fmt.Println("gRPC server started on :50051")
	return grpcServer.Serve(lis)
}
