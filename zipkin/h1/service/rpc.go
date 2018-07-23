package service

import (
	"context"

	s1 "github.com/JREAMLU/study/zipkin/s1/proto"
)

// GetA get a
func GetA(ctx context.Context) (string, error) {
	resp, err := s1Client.AHello(ctx, &s1.AHelloRequest{
		Name: "LU",
	})

	if err != nil {
		return "", err
	}

	return resp.Greeting, nil
}
