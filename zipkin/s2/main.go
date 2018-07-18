package main

import (
	"time"

	"github.com/JREAMLU/study/zipkin/s2/controller"
	proto "github.com/JREAMLU/study/zipkin/s2/proto"
	"github.com/JREAMLU/study/zipkin/s2/service"
	micro "github.com/micro/go-micro"
)

func main() {
	RunMicroService()
}

// RunMicroService run micro service
func RunMicroService() {
	// Create a new service. Optionally include some options here.
	ms := micro.NewService(
		micro.Name("go.micro.srv.s2"),
		micro.Version("v1"),
	)

	// Init will parse the command line flags.
	ms.Init(
		micro.RegisterTTL(1*time.Second),
		micro.RegisterInterval(1*time.Second),
	)

	// Register handler
	proto.RegisterS2Handler(ms.Server(), controller.NewS2Handler())

	service.InitMicroClient(ms.Client())

	// Run the server
	if err := ms.Run(); err != nil {
		panic(err)
	}
}
