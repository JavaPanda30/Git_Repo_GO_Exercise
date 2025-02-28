package main

import (
	"context"

	pb "example.com/aman/basic-grpc-unary/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hi this is a message the first initialized by server side then called using rpc in client side as if this function is there with the client.",
	}, nil
}
