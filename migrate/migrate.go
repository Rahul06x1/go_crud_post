package main

import (
	"github.com/Rahul06x1/go_crud/initializers"
	"github.com/Rahul06x1/go_crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}