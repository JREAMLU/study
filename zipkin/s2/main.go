package main

import (
	"github.com/JREAMLU/study/zipkin/s2/config"
	"github.com/JREAMLU/study/zipkin/s2/controller"
	proto "github.com/JREAMLU/study/zipkin/s2/proto"
	"github.com/JREAMLU/study/zipkin/s2/service"

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
func RunMicroService(conf *config.S2Config) {
	ms := util.NewMicroService(conf.Config)

	// Register handler
	proto.RegisterS2Handler(ms.Server(), controller.NewS2Handler())

	// init client
	service.InitMicroClient(ms.Client())

	// Run the server
	if err := ms.Run(); err != nil {
		panic(err)
	}
}
