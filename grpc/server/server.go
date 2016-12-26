package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/JREAMLU/study/grpc/fly"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) DriverPlane(ctx context.Context, in *pb.FlyRequest) (*pb.FlyReply, error) {
	return &pb.FlyReply{Name: "Driver Version " + in.Version}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFlyToSkyServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
