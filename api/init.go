package api

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func Run() {
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	// 使用 CORS 中间件
	r.Use(CORSMiddleware())

	group := r.Group("/project")
	group.POST("/create", createProject)

	{
		functionGroup := group.Group("/function")
		functionGroup.POST("/create", createProjectFunctions)
		functionGroup.POST("/sure", makeSureProjectFunctions)
	}
	{
		modelGroup := group.Group("/model")
		modelGroup.POST("/create", createProjectModels)
		modelGroup.POST("/sure", makeSureProjectModels)
	}
	group.POST("/code", code)

	if err := r.Run(":18081"); err != nil {
		panic("failed to start server")
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
