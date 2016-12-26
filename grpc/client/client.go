package main

import (
	"context"
	"log"

	pb "github.com/JREAMLU/study/grpc/fly"
	"google.golang.org/grpc"
)

const (
	address        = "localhost:50051"
	defaultVersion = "V1.1.1"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewFlyToSkyClient(conn)

	r, err := c.DriverPlane(context.Background(), &pb.FlyRequest{Version: defaultVersion})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Version: %s", r.Name)
}
