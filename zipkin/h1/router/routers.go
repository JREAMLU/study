package router

import (
	"github.com/JREAMLU/study/zipkin/h1/config"
	"github.com/JREAMLU/study/zipkin/h1/controller"
	"github.com/JREAMLU/study/zipkin/h1/middleware"
	"github.com/gin-gonic/gin"
)

// GetRouters init router
func GetRouters(router *gin.Engine, conf *config.HelloConfig) *gin.Engine {
	// hello world
	router.GET("/h1", middleware.Middle(), controller.NewHelloController(conf).World)
	router.POST("/h1", middleware.Middle(), controller.NewHelloController(conf).WorldP)

	return router
}
