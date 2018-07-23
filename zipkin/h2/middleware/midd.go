package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Middle middleware
func Middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "client_request")
		// 取出所有的注册的函数都执行一遍, 然后再回到本函数中
		c.Next()
		fmt.Println("before middleware")
	}
}
