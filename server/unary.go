package main

import (
	"context"

	pb "github.com/MehulxBuilds/grpc-basics/proto"
)
func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}