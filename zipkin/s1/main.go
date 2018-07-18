package main

import (
	"time"

	"github.com/JREAMLU/study/zipkin/s1/controller"
	proto "github.com/JREAMLU/study/zipkin/s1/proto"
	"github.com/JREAMLU/study/zipkin/s1/service"
	"github.com/JREAMLU/study/zipkin/util"
	micro "github.com/micro/go-micro"

	// brokerKafka "github.com/micro/go-plugins/broker/kafka"
	clientGrpc "github.com/micro/go-plugins/client/grpc"
	registerConsul "github.com/micro/go-plugins/registry/consul"
	serverGrpc "github.com/micro/go-plugins/server/grpc"
	transportGrpc "github.com/micro/go-plugins/transport/grpc"
)

func main() {
	RunMicroService()
}

// RunMicroService run micro service
func RunMicroService() {
	// Create a new service. Optionally include some options here.
	ms := micro.NewService(
		micro.Client(clientGrpc.NewClient()),
		micro.Server(serverGrpc.NewServer()),
		// micro.Broker(brokerKafka.NewBroker(
		// 	broker.Option(func(opt *broker.Options) {
		// 		opt.Addrs = []string{"10.200.119.128:9092"}
		// 	}),
		// )),
		micro.Registry(registerConsul.NewRegistry()),
		micro.Transport(transportGrpc.NewTransport()),
		micro.Name("go.micro.srv.s1"),
		micro.Version("v1"),
		micro.WrapClient(util.TraceWrapperClent),
		micro.WrapHandler(util.TraceWrapperHandler),
	)

	// Init will parse the command line flags.
	ms.Init(
		micro.RegisterTTL(1*time.Second),
		micro.RegisterInterval(1*time.Second),
	)

	// Register handler
	proto.RegisterS1Handler(ms.Server(), controller.NewS1Handler())

	service.InitMicroClient(ms.Client())

	// Run the server
	if err := ms.Run(); err != nil {
		panic(err)
	}
}
