package service

import (
	s3PB "github.com/JREAMLU/study/zipkin/s3/proto"
	microClient "github.com/micro/go-micro/client"
)

const (
	// S3 service name
	S3 = "go.micro.srv.s3"
)

var (
	s3Client s3PB.S3Service
)

// InitMicroClient init micro client
func InitMicroClient(c microClient.Client) {
	s3Client = s3PB.S3ServiceClient(S3, c)
}
