package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vvatelot/go-upload-files/handlers"
	"github.com/vvatelot/go-upload-files/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	services.AuthorizedUserIds = strings.Split(os.Getenv("AUTHORIZED_USERID"), ",")

	router := gin.Default()
	router.LoadHTMLFiles("templates/index.tmpl", "templates/forbidden.tmpl")
	router.Static("/public", "./public")
	router.GET("/", handlers.HandleHome)
	router.POST("/upload", handlers.HandleUpload)
	router.Run(":8080")
}
