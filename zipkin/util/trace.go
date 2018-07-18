package util

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
)

// TraceWrapper trace wrapper
type traceWrapper struct {
	client.Client
}

func (t *traceWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[wrapper] client request service: %s method: %s\n", req.Service(), req.Method())
	return t.Client.Call(ctx, req, rsp)
}

// TraceWrapperClent trace wraper client
func TraceWrapperClent(c client.Client) client.Client {
	return &traceWrapper{c}
}

// TraceWrapperHandler trace wraper handler
func TraceWrapperHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		fmt.Printf("[wrapper] server request: %v\n", req.Method())
		return fn(ctx, req, resp)
	}
}
