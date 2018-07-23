package main

import (
	"github.com/JREAMLU/study/zipkin/h2/config"
	"github.com/JREAMLU/study/zipkin/h2/router"
	"github.com/JREAMLU/study/zipkin/h2/service"

	"github.com/gin-gonic/gin"
	microClient "github.com/micro/go-plugins/client/grpc"
)

func main() {

	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	service.InitMicroClient(microClient.NewClient())

	g := gin.New()
	g.Use(gin.Recovery(), gin.Logger())
	g = router.GetRouters(g, conf)
	g.Run(":8002")
}
