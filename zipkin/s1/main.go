package main

import (
	"github.com/JREAMLU/study/zipkin/s1/config"
	"github.com/JREAMLU/study/zipkin/s1/controller"
	proto "github.com/JREAMLU/study/zipkin/s1/proto"
	"github.com/JREAMLU/study/zipkin/s1/service"

	"github.com/JREAMLU/j-kit/go-micro/util"
)

func main() {
	// load config
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	RunMicroService(conf)
}

// RunMicroService run micro service
func RunMicroService(conf *config.S1Config) {
	ms := util.NewMicroService(conf.Config)

	// Register handler
	proto.RegisterS1Handler(ms.Server(), controller.NewS1Handler())

	// init client
	service.InitMicroClient(ms.Client())

	// Run the server
	if err := ms.Run(); err != nil {
		panic(err)
	}
}
