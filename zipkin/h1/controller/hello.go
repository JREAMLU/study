package controller

import (
	"fmt"
	"net/http"

	"github.com/JREAMLU/study/zipkin/h1/config"
	"github.com/JREAMLU/study/zipkin/h1/service"
	"github.com/gin-gonic/gin"
)

// HelloController hello controller
type HelloController struct {
	Controller
}

// NewHelloController new hello
func NewHelloController(conf *config.HelloConfig) *HelloController {
	return &HelloController{
		Controller{
			config: conf,
		},
	}
}

// World world
func (h *HelloController) World(c *gin.Context) {
	ctx := c.Request.Context()
	err := service.Geth2(ctx)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// s, err := service.GetA(ctx)
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	return
	// }

	request := c.MustGet("request").(string)
	c.String(http.StatusOK, "Hello World %v, %v", request, "")
}

// WorldP World Post
func (h *HelloController) WorldP(c *gin.Context) {
	ctx := c.Request.Context()

	// err := service.GetH1(ctx)
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	return
	// }

	err := service.Geth2P(ctx)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// s, err := service.GetA(ctx)
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	return
	// }

	request := c.MustGet("request").(string)
	c.String(http.StatusOK, "Hello World %v, %v", request, "")
}
