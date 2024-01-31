package controllers

import (
	"github.com/Rahul06x1/go_crud/initializers"
	"github.com/Rahul06x1/go_crud/models"

	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

func router() *gin.Engine {
	router := gin.Default()

	publicRoutes := router.Group("/api")
	// protectedRoutes.Use(middleware.JWTAuthMiddleware())

	publicRoutes.GET("/posts", PostsIndex)
	publicRoutes.GET("/posts/:id", PostsShow)
	publicRoutes.POST("/posts", PostCreate)
	publicRoutes.PATCH("/posts/:id", PostsUpdate)
	publicRoutes.DELETE("/posts/:id", PostsDelete)

	return router
}

func setup() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Post{})
}

func teardown() {
	migrator := initializers.DB.Migrator()
	migrator.DropTable(&models.Post{})
}
