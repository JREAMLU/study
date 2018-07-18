package main

import (
	"time"

	"github.com/JREAMLU/study/zipkin/s3/controller"
	proto "github.com/JREAMLU/study/zipkin/s3/proto"
	"github.com/JREAMLU/study/zipkin/s3/service"
	micro "github.com/micro/go-micro"
)

func main() {
	RunMicroService()
}

// RunMicroService run micro service
func RunMicroService() {
	// Create a new service. Optionally include some options here.
	ms := micro.NewService(
		micro.Name("go.micro.srv.s3"),
		micro.Version("v1"),
	)

	// Init will parse the command line flags.
	ms.Init(
		micro.RegisterTTL(1*time.Second),
		micro.RegisterInterval(1*time.Second),
	)

	// Register handler
	proto.RegisterS3Handler(ms.Server(), controller.NewS3Handler())

	service.InitMicroClient(ms.Client())

	// Run the server
	if err := ms.Run(); err != nil {
		panic(err)
	}
}
