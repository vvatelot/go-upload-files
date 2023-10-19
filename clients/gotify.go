package clients

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func SendGotifyNotification(nbFiles int, userName string) {
	formData := url.Values{
		"title":   {"Nouveaux fichiers sur le serveur"},
		"message": {fmt.Sprintf("%d fichiers ont été envoyés par %s", nbFiles, userName)},
	}

	_, err := http.PostForm(os.Getenv("GOTIFY_URL")+"/message?token="+os.Getenv("GOTIFY_TOKEN"), formData)
	if err != nil {
		log.Fatalln(err)
	}
}
