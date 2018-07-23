package router

import (
	"github.com/JREAMLU/study/zipkin/h2/config"
	"github.com/JREAMLU/study/zipkin/h2/controller"
	"github.com/JREAMLU/study/zipkin/h2/middleware"
	"github.com/gin-gonic/gin"
)

// GetRouters init router
func GetRouters(router *gin.Engine, conf *config.HelloConfig) *gin.Engine {
	// hello world
	router.GET("/h2", middleware.Middle(), controller.NewHelloController(conf).World)

	return router
}
