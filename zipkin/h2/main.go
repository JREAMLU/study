package main

import (
	"github.com/JREAMLU/j-kit/http"
	"github.com/JREAMLU/study/zipkin/h2/config"
	"github.com/JREAMLU/study/zipkin/h2/router"
	"github.com/JREAMLU/study/zipkin/h2/service"
)

func main() {
	// load config
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	RunHTTPService(conf)
}

// RunHTTPService run http service
func RunHTTPService(conf *config.HelloConfig) {
	ms, g, t := http.NewHTTPService(conf.Config)

	// init micro client
	service.InitMicroClient(ms.Client())

	// init http client
	service.InitHTTPClient(t)

	g = router.GetRouters(g, conf)
	g.Run(conf.Web.URL)
}
