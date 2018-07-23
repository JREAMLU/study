package main

import (
	"github.com/JREAMLU/study/zipkin/h3/config"
	"github.com/JREAMLU/study/zipkin/h3/router"

	"github.com/gin-gonic/gin"
)

func main() {

	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	g := gin.New()
	g.Use(gin.Recovery(), gin.Logger())
	g = router.GetRouters(g, conf)
	g.Run(":8003")
}
