package service

import (
	"context"

	s2 "github.com/JREAMLU/study/zipkin/s2/proto"
	s3 "github.com/JREAMLU/study/zipkin/s3/proto"
)

// GetB get b
func GetB(ctx context.Context) (string, error) {
	m := make(map[string]string)
	m["x"] = "y"

	resp, err := s2Client.BHello(ctx, &s2.BHelloRequest{
		Name:   "Iverson",
		Rid:    []string{"abc", "edf"},
		Extras: m,
	})

	if err != nil {
		return "", err
	}

	return resp.Greeting, nil
}

// GetC get c
func GetC(ctx context.Context) (string, error) {
	resp, err := s3Client.CHello(ctx, &s3.CHelloRequest{
		Name: "KD",
	})

	if err != nil {
		return "", err
	}

	return resp.Greeting, nil
}
