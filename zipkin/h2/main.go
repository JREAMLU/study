package main

import (
	"github.com/JREAMLU/j-kit/go-micro/util"
	"github.com/JREAMLU/study/zipkin/h2/config"
	"github.com/JREAMLU/study/zipkin/h2/router"
	"github.com/JREAMLU/study/zipkin/h2/service"

	"github.com/gin-gonic/gin"
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

	service.InitHTTPClient(t)

	g := gin.New()
	g.Use(gin.Recovery(), gin.Logger())
	g = router.GetRouters(g, conf)
	g.Run(":8002")
}
