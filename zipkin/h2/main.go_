package main

import (
	"time"

	"github.com/JREAMLU/j-kit/go-micro/trace/opentracing"
	"github.com/JREAMLU/j-kit/go-micro/util"
	"github.com/JREAMLU/j-kit/http"
	"github.com/JREAMLU/study/zipkin/h2/config"
	"github.com/JREAMLU/study/zipkin/h2/router"
	"github.com/JREAMLU/study/zipkin/h2/service"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	clientGrpc "github.com/micro/go-plugins/client/grpc"
	registerConsul "github.com/micro/go-plugins/registry/consul"
	transportGrpc "github.com/micro/go-plugins/transport/grpc"
	// microClient "github.com/micro/go-plugins/client/grpc"
)

func main() {
	t, err := util.NewTrace("go.http.srv.h2", "v1", []string{"10.200.119.128:9092", "10.200.119.129:9092", "10.200.119.130:9092"}, "web_log_get")
	if err != nil {
		panic(err)
	}

	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	ms := micro.NewService(
		micro.Client(clientGrpc.NewClient()),
		micro.Registry(registerConsul.NewRegistry()),
		micro.Transport(transportGrpc.NewTransport()),
		micro.WrapClient(opentracing.NewClientWrapper(t)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
	)
	ms.Init(
		micro.RegisterTTL(1*time.Second),
		micro.RegisterInterval(1*time.Second),
	)
	service.InitMicroClient(ms.Client())
	// service.InitMicroClient(microClient.NewClient())
	service.InitHTTPClient(t)

	g := gin.New()
	g.Use(gin.Recovery(), gin.Logger(), http.HandlerHTTPRequestGin(t, "go.http.srv.h2"))

	g = router.GetRouters(g, conf)
	g.Run(":8002")
}
