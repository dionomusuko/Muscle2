package main

import (
	"github.com/dionomusuko/muscle2/controller"
	"github.com/dionomusuko/muscle2/db"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := db.NewDB()
	defer conn.Close()

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", controller.GetUsers)
			users.GET("/:id", controller.ShowUser)
			users.POST("/", controller.CreateUser)
			users.PUT("/:id", controller.UpdateUser)
			users.DELETE("/:id", controller.DeleteUser)
		}
		tasks := v1.Group("/tasks")
		{
			tasks.GET("/", controller.GetTasks)
			tasks.GET("/:id", controller.ShowTask)
			tasks.POST("/", controller.CreateTask)
			tasks.PUT("/:id", controller.UpdateTask)
			tasks.DELETE("/:id", controller.DeleteTask)
		}
	}
	router.Run(":8080")
}
