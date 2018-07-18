package controller

import (
	"context"

	proto "github.com/JREAMLU/study/zipkin/s2/proto"
	"github.com/JREAMLU/study/zipkin/s2/service"
)

// S2 s2
type S2 struct{}

// NewS2Handler new greeter
func NewS2Handler() *S2 {
	return &S2{}
}

// BHello h w
func (s *S2) BHello(ctx context.Context, req *proto.BHelloRequest, resp *proto.BHelloResponse) error {
	s3, err := service.GetC(ctx)
	if err != nil {
		return err
	}

	resp.Greeting = "S2: BHello " + req.Name + " " + s3
	return nil
}
