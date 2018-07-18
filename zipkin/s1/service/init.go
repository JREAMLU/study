package service

import (
	s2PB "github.com/JREAMLU/study/zipkin/s2/proto"
	s3PB "github.com/JREAMLU/study/zipkin/s3/proto"
	microClient "github.com/micro/go-micro/client"
)

const (
	// S2 service name
	S2 = "go.micro.srv.s2"
	// S3 service name
	S3 = "go.micro.srv.s3"
)

var (
	s2Client s2PB.S2Service
	s3Client s3PB.S3Service
)

// InitMicroClient init micro client
func InitMicroClient(c microClient.Client) {
	s2Client = s2PB.S2ServiceClient(S2, c)
	s3Client = s3PB.S3ServiceClient(S3, c)
}
