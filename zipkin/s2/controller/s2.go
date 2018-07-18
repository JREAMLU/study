package controller

import (
	"context"

	proto "github.com/JREAMLU/study/zipkin/s2/proto"
)

// S2 s2
type S2 struct{}

// NewS2Handler new greeter
func NewS2Handler() *S2 {
	return &S2{}
}

// BHello h w
func (s *S2) BHello(ctx context.Context, req *proto.BHelloRequest, resp *proto.BHelloResponse) error {
	resp.Greeting = "BHello " + req.Name
	return nil
}
