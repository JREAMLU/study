package controller

import (
	"context"

	proto "github.com/JREAMLU/study/zipkin/s3/proto"
)

// S3 s3
type S3 struct{}

// NewS3Handler new greeter
func NewS3Handler() *S3 {
	return &S3{}
}

// CHello h w
func (s *S3) CHello(ctx context.Context, req *proto.CHelloRequest, resp *proto.CHelloResponse) error {
	resp.Greeting = "BHello " + req.Name
	return nil
}
