package controller

import (
	"context"
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
	s, err := service.GetA(context.Background())
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	request := c.MustGet("request").(string)
	c.String(http.StatusOK, "Hello World %v, %v", request, s)
}
