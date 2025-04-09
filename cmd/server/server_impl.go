package main

import (
	"context"
	"log"

	pb "golang-grpc-example/internal/greet"
)

// server is used to implement greet.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements greet.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func NewServer() *server {
	return &server{}
}
