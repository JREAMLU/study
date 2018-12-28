package controller

import (
	"context"

	pb "github.com/JREAMLU/study/go-micro/server/proto"
)

// Greeter greeter
type Greeter struct{}

// NewGreeterHandler new greeter
func NewGreeterHandler() *Greeter {
	return &Greeter{}
}

// Hello h w
func (g *Greeter) Hello(ctx context.Context, req *pb.HelloRequest, resp *pb.HelloResponse) error {
	resp.Greeting = "Hello " + req.Name
	return nil
}
