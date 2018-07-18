package service

import (
	"context"

	s2 "github.com/JREAMLU/study/zipkin/s2/proto"
)

// GetB get b
func GetB(ctx context.Context) (string, error) {
	resp, err := s2Client.BHello(ctx, &s2.BHelloRequest{
		Name: "Iverson",
	})

	if err != nil {
		return "", err
	}

	return resp.Greeting, nil
}
