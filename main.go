package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		userId := c.Query("userid")

		if userId != "vvatel22" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Non autorisé",
			})
			return
		}

		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("get form err: %s", err.Error()),
			})
			return
		}
		files := form.File["files"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("upload file err: %s", err.Error()),
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Uploaded successfully %d files.", len(files)),
		})
	})
	router.Run(":8080")
}
