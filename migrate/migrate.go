package main

import (
	"github.com/ahadalichowdhury/go_crud/initializers"
	"github.com/ahadalichowdhury/go_crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	 // Migrate the schema
  initializers.DB.AutoMigrate(&models.Post{})
}