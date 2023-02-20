package controllers

import (
	"github.com/ahadalichowdhury/go_crud/initializers"
	"github.com/ahadalichowdhury/go_crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	//get data from the request body

	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	//create a post

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostsRead(c *gin.Context) {
	//get the post 
	var posts []models.Post
	initializers.DB.Find(&posts)

	//response
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostRead(c *gin.Context) {

	//get id
	id:= c.Param("id")

		//get the post 
	var post models.Post
	initializers.DB.First(&post, id)

	//response
	c.JSON(200, gin.H{
		"post": post,
	})
	  
}


func PostUpdate(c *gin.Context){
	//get the id from the URL
		id:= c.Param("id")


	//get the data from the request body
	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	//find the post where we need to update
		var post models.Post
	initializers.DB.First(&post, id)

	//update it

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//respond with it
	c.JSON(200, gin.H{
		"post": post,
	})

}


func PostDelete(c *gin.Context) {

	//get id
	id:= c.Param("id")

	//delete the post
	initializers.DB.Delete(&models.Post{}, id)

	//response
	c.Status(200)
	  
}