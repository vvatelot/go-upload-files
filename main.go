package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var authorizedUserIds []string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	authorizedUserIds = strings.Split(os.Getenv("AUTHORIZED_USERID"), ",")

	router := gin.Default()
	router.LoadHTMLFiles("templates/index.tmpl", "templates/forbidden.tmpl")
	router.Static("/public", "./public")
	router.GET("/", handleHome)
	router.POST("/upload", handleUpload)
	router.Run(":8080")
}

func handleHome(c *gin.Context) {
	userId := c.Query("userid")

	if !checkUserid(userId) {
		c.HTML(http.StatusForbidden, "forbidden.tmpl", gin.H{})
	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"userName": getUserName(userId),
		})
	}
}

func handleUpload(c *gin.Context) {
	userId := c.Query("userid")

	if !checkUserid(userId) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "⛔ Non autorisé",
		})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("❌ Erreur de formulaire : %s", err.Error()),
		})
		return
	}

	userName := getUserName(userId)
	files := form.File["files"]

	for _, file := range files {
		filename := fmt.Sprintf("%s%s %s%s", os.Getenv("TARGET_FOLDER"), userName, time.Now().String(), filepath.Ext(file.Filename))
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("❌ Fichier(s) non envoyé(s) : %s", err.Error()),
			})
			return
		}
	}

	sendGotifyNotification(len(files), userName)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("🎉 %d fichier(s) envoyé(s) avec succès", len(files)),
	})
}

func checkUserid(userId string) bool {
	for _, auhorizedUserid := range authorizedUserIds {
		if userId == auhorizedUserid {
			return true
		}
	}
	return false
}

func sendGotifyNotification(nbFiles int, userName string) {
	formData := url.Values{
		"title":   {"Nouveaux fichiers sur le serveur"},
		"message": {fmt.Sprintf("%d fichiers ont été envoyés par %s", nbFiles, userName)},
	}

	_, err := http.PostForm(os.Getenv("GOTIFY_URL")+"/message?token="+os.Getenv("GOTIFY_TOKEN"), formData)
	if err != nil {
		log.Fatalln(err)
	}
}

func getUserName(userId string) string {
	userIdParts := strings.Split(userId, "|")
	userName := userIdParts[0]

	return strings.Replace(userName, "_", " ", -1)
}
