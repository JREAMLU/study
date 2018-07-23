package controller

import (
	"net/http"

	"github.com/JREAMLU/study/zipkin/h3/config"
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
	request := c.MustGet("request").(string)
	c.String(http.StatusOK, "Hello World 3 %v", request)
}
