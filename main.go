package main

import (
	"github.com/ahadalichowdhury/go_crud/controllers"
	"github.com/ahadalichowdhury/go_crud/initializers"
	"github.com/gin-gonic/gin"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostsRead)
	r.GET("posts/:id", controllers.PostRead)
	r.PUT("posts/:id", controllers.PostUpdate)
	r.DELETE("posts/:id", controllers.PostDelete)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}