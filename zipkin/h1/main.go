package main

import (
	"github.com/JREAMLU/j-kit/go-micro/util"
	"github.com/JREAMLU/j-kit/http"
	"github.com/JREAMLU/study/zipkin/h1/config"
	"github.com/JREAMLU/study/zipkin/h1/router"
	"github.com/JREAMLU/study/zipkin/h1/service"

	"github.com/gin-gonic/gin"
)

func main() {
}

// RunHTTPService run http service
func RunHTTPService(conf *config.HelloConfig) {
	ms, t := util.NewHTTPService(conf.Config)

	// init micro client
	service.InitMicroClient(ms.Client())

	// init http client
	service.InitHTTPClient(t)

	g := gin.New()
	g.Use(
		gin.Recovery(),
		gin.Logger(),
		http.HandlerHTTPRequestGin(t, conf.Service.Name),
	)

	g = router.GetRouters(g, conf)
	g.Run(conf.Web.URL)
}
