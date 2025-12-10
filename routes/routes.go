package routes

import (
	"my-task-manager/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/tasks", CreateTask)
	authenticated.PUT("/tasks/:id", UpdateTask)

	server.GET("/tasks", GetTasks)
	server.POST("/login", login)
	server.POST("/signup", signUp)
}
