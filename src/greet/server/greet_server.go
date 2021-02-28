package main

import (
	"context"
	"log"
	"net"

	pb "greeting"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
}

func (s *server) Greet(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.GreetingResponse{Message: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})
	log.Print("Greeting server successfully started...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
