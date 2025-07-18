package router

import (
	"github.com/buzyka/about_go/06_gin_project/internal/infrastructure/http/handler"
	"github.com/buzyka/about_go/06_gin_project/internal/infrastructure/http/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.AuthMiddleware())
	apiGroup.POST("/create-full-name", handler.CreateFullName)

	return r
}
