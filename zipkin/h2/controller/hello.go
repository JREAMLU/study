package controller

import (
	"fmt"
	"net/http"

	"github.com/JREAMLU/study/zipkin/h2/config"
	"github.com/JREAMLU/study/zipkin/h2/service"
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
	err := service.Geth3(ctx)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	request := c.MustGet("request").(string)
	c.String(http.StatusOK, "Hello World 2 %v", request)
}

// WorldP world post
func (h *HelloController) WorldP(c *gin.Context) {
	ctx := c.Request.Context()
	err := service.Geth3(ctx)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	request := c.MustGet("request").(string)
	c.String(http.StatusOK, "Hello World 2 %v", request)
}
