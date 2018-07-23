package router

import (
	"github.com/JREAMLU/study/zipkin/h3/config"
	"github.com/JREAMLU/study/zipkin/h3/controller"
	"github.com/JREAMLU/study/zipkin/h3/middleware"
	"github.com/gin-gonic/gin"
)

// GetRouters init router
func GetRouters(router *gin.Engine, conf *config.HelloConfig) *gin.Engine {
	// hello world
	router.GET("/h3", middleware.Middle(), controller.NewHelloController(conf).World)

	return router
}
