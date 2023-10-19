package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vvatelot/go-upload-files/clients"
	"github.com/vvatelot/go-upload-files/services"
)

func HandleHome(c *gin.Context) {
	userId := c.Query("userid")

	if !services.CheckUserid(userId) {
		c.HTML(http.StatusForbidden, "forbidden.tmpl", gin.H{})
	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"userName": services.GetUserName(userId),
		})
	}
}

func HandleUpload(c *gin.Context) {
	userId := c.Query("userid")

	if !services.CheckUserid(userId) {
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

	userName := services.GetUserName(userId)
	files := form.File["files"]

	for _, file := range files {
		t := time.Now()
		filename := fmt.Sprintf("%s%d-%02d-%02dT%02d:%02d:%02d:%d-%s%s", os.Getenv("TARGET_FOLDER"), t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), userName, filepath.Ext(file.Filename))
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("❌ Fichier(s) non envoyé(s) : %s", err.Error()),
			})
			return
		}
	}

	if os.Getenv("GOTIFY_URL") != "" && os.Getenv("GOTIFY_TOKEN") != "" {
		clients.SendGotifyNotification(len(files), userName)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("🎉 %d fichier(s) envoyé(s) avec succès", len(files)),
	})
}
