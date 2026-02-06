package main

import (
	"fmt"
	"goorm/controllers"
	initiaizers "goorm/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initiaizers.LoadEnvVariables()
	initiaizers.ConnectDB()
}

func main() {
	fmt.Println("hello")
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/post", controllers.Postindex)
	r.GET("/post/:id", controllers.PostShow)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
}
