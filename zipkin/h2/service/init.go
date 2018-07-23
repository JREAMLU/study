package service

import (
	"github.com/JREAMLU/j-kit/http"

	opentracing "github.com/opentracing/opentracing-go"
)

var (
	httpClient *http.Requests
)

// InitHTTPClient init http client
func InitHTTPClient(tracer opentracing.Tracer) {
	httpClient = http.NewRequests(tracer)
}
