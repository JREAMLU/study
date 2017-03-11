package main

import (
	"log"
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-os/trace"
	_ "github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/go-plugins/trace/zipkin"

	"github.com/JREAMLU/study/micro/service/proto"
	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Nickname = "Hello " + req.Name
	return nil
}

func main() {
	// t := trace.NewTrace()

	// srv := &registry.Service{Name: "greeter", Version: "v1"}

	service := grpc.NewService(
		micro.Name("greeter"),
		micro.Version("v1"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		// micro.WrapClient(trace.ClientWrapper(t, srv)),
		// micro.WrapHandler(trace.HandlerWrapper(t, srv)),
	)

	service.Init(Trace("greeter", "v1"))

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func Trace(serviceNmae, version string) micro.Option {
	return func(opt *micro.Options) {
		t := zipkin.NewTrace(trace.Client(opt.Client), trace.Collectors("172.16.9.4:9092"))
		srv := &registry.Service{Name: serviceNmae, Version: version}
		micro.WrapClient(trace.ClientWrapper(t, srv))(opt)
		micro.WrapHandler(trace.HandlerWrapper(t, srv))(opt)

		// t := trace.NewTrace(trace.Client(opt.Client))
		// srv := &registry.Service{Name: serviceNmae, Version: version}
		// micro.WrapClient(trace.ClientWrapper(t, srv))(opt)
		// micro.WrapHandler(trace.HandlerWrapper(t, srv))(opt)
	}
}
