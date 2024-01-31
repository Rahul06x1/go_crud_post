package controllers

import (
	"github.com/Rahul06x1/go_crud/initializers"
	"github.com/Rahul06x1/go_crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// Get post off request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(201, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id") 

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off url 
	id := c.Param("id")

	// Get data off req body 
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the post were updating 
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it 
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Respond with it 
	c.JSON(201, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off url 
	id := c.Param("id")

	// Delete the post 
	initializers.DB.Delete(&models.Post{}, id) // Normal delete
	// initializers.DB.Unscoped().Delete(&models.Post{}, id) // Hard delete (deletes from db)

	// Respond 
	c.Status(200)
}