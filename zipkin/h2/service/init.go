package service

import (
	"github.com/JREAMLU/j-kit/http"

	s1PB "github.com/JREAMLU/study/zipkin/s1/proto"

	microClient "github.com/micro/go-micro/client"
	opentracing "github.com/opentracing/opentracing-go"
)

const (
	// S1 service name
	S1 = "go.micro.srv.s1"
)

var (
	s1Client   s1PB.S1Service
	httpClient *http.Requests
)

// InitMicroClient init micro client
func InitMicroClient(c microClient.Client) {
	s1Client = s1PB.S1ServiceClient(S1, c)
}

// InitHTTPClient init http client
func InitHTTPClient(tracer opentracing.Tracer) {
	httpClient = http.NewRequests(tracer)
}
