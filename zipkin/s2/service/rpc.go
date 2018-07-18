package service

import (
	"context"

	s3 "github.com/JREAMLU/study/zipkin/s3/proto"
)

// GetC get b
func GetC(ctx context.Context) (string, error) {
	resp, err := s3Client.CHello(ctx, &s3.CHelloRequest{
		Name: "Curry",
	})

	if err != nil {
		return "", err
	}

	return resp.Greeting, nil
}
