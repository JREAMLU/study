package main

import (
	"log"
	"time"

	"github.com/JREAMLU/study/micro/service/proto"
	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-os/trace"
	"golang.org/x/net/context"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Nickname = "Hello " + req.Name
	return nil
}

func main() {
	t := trace.NewTrace()

	srv := &registry.Service{Name: "greeter", Version: "v1"}

	service := grpc.NewService(
		micro.Name("greeter"),
		micro.Version("v1"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.WrapClient(trace.ClientWrapper(t, srv)),
		micro.WrapHandler(trace.HandlerWrapper(t, srv)),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
