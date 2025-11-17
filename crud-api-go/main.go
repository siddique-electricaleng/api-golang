package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.POST("/posts", controllers.PostsCreate)
	router.PUT("/posts/:id", controllers.PostsUpdate)

	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)

	router.Run()
}
