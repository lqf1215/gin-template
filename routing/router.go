package routing

import (
	"gin-template/intercept"
	"github.com/gin-gonic/gin"
)

func Setup(e *gin.Engine) {
	g := e.Group("/api")
	g.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	gAuth := g.Use(intercept.AuthApp())
	gAuth.GET("/login", func(c *gin.Context) {
		c.String(200, "123")
	})
}
