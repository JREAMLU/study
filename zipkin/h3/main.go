package main

import (
	"github.com/JREAMLU/j-kit/http"
	"github.com/JREAMLU/study/zipkin/h3/config"
	"github.com/JREAMLU/study/zipkin/h3/router"
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
	_, g, _ := http.NewHTTPService(conf.Config)

	g = router.GetRouters(g, conf)
	g.Run(conf.Web.URL)
}
