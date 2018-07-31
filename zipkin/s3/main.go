package main

import (
	"github.com/JREAMLU/study/zipkin/s3/config"
	"github.com/JREAMLU/study/zipkin/s3/controller"
	proto "github.com/JREAMLU/study/zipkin/s3/proto"
	"github.com/JREAMLU/study/zipkin/s3/service"

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
func RunMicroService(conf *config.S3Config) {
	ms := util.NewMicroService(conf.Config)

	// Register handler
	proto.RegisterS3Handler(ms.Server(), controller.NewS3Handler())

	// init client
	service.InitMicroClient(ms.Client())

	// Run the server
	if err := ms.Run(); err != nil {
		panic(err)
	}
}
