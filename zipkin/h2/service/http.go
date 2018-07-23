package service

import (
	"context"
)

// Geth3 h3
func Geth3(ctx context.Context) error {
	_, err := httpClient.RequestCURL(ctx,
		"GET", "http://127.0.0.1:8003/h3",
		map[string]string{
			"Content-Type": "application/json;charset=UTF-8;",
		},
		`{"name":"K3","age":26}`,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
