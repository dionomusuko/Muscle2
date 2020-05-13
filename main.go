package main

import (
	Controller "github.com/dionomusuko/muscle2/controller"
	"github.com/dionomusuko/muscle2/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func main() {
	conn := db.NewDB()
	defer conn.Close()

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", Controller.GetUsers)
			users.GET("/:id", Controller.ShowUser)
			users.POST("/", Controller.CreateUser)
			users.PUT("/:id", Controller.UpdateUser)
			users.DELETE("/:id", Controller.DeleteUser)
		}
		tasks := v1.Group("/tasks")
		{
			tasks.GET("/", Controller.GetTasks)
			tasks.GET("/:id", Controller.ShowTask)
			//tasks.POST("/", Controller.CreateTask)
			tasks.PUT("/:id", Controller.UpdateTask)
			tasks.DELETE("/:id", Controller.DeleteTask)
		}
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
	}))
	router.Run(":8080")
}
