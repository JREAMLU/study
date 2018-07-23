package service

import (
	s1PB "github.com/JREAMLU/study/zipkin/s1/proto"
	microClient "github.com/micro/go-micro/client"
)

const (
	// S1 service name
	S1 = "go.micro.srv.s1"
)

var (
	s1Client s1PB.S1Service
)

// InitMicroClient init micro client
func InitMicroClient(c microClient.Client) {
	s1Client = s1PB.S1ServiceClient(S1, c)
}
