package controller

import (
	"context"

	proto "github.com/JREAMLU/study/zipkin/s1/proto"
)

// S1 s1
type S1 struct{}

// NewS1Handler new greeter
func NewS1Handler() *S1 {
	return &S1{}
}

// AHello h w
func (s *S1) AHello(ctx context.Context, req *proto.AHelloRequest, resp *proto.AHelloResponse) error {
	resp.Greeting = "AHello " + req.Name
	return nil
}
