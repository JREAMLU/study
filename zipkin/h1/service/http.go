package service

import (
	"context"
)

// Geth2 h2
func Geth2(ctx context.Context) error {
	_, err := httpClient.RequestCURL(ctx,
		"GET", "http://127.0.0.1:8002/h2",
		map[string]string{
			"Content-Type": "application/json;charset=UTF-8;",
		},
		`{"name":"KII","age":24}`,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

// Geth2P h2 post
func Geth2P(ctx context.Context) error {
	_, err := httpClient.RequestCURL(ctx,
		"POST", "http://127.0.0.1:8002/h2",
		map[string]string{
			"Content-Type": "application/json;charset=UTF-8;",
		},
		`{"name":"KII","age":24}`,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
