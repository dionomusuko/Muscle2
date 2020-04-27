package main

import (
	"github.com/gin-gonic/gin"
	"golang/go-api/db"
)

func main() {
	db.Init()
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", UserController.GetUsers)
			users.GET("/:id", UserController.ShowUser)
			users.POST("/", UserController.CreateUser)
			users.PUT("/:id", UserController.UpdateUser)
			users.DELETE("/:id", UserController.Deleteuser)
		}
		tasks := v1.Group("/tasks")
		{
			tasks.GET("/", TaskController.GetTasks)
			tasks.GET("/:id", TaskController.ShowUser)
			tasks.POST("/", TaskController.CreateTask)
			tasks.PUT("/:id", TaskController.UpdateTask)
			tasks.DELETE("/:id", TaskController.DeleteTask)
		}
	}

}
