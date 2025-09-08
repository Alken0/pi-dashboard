package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Feature struct {
	Image       string
	Title       string
	Description string
	Link        string
}

func HomePage(c *gin.Context) {
	features := []Feature{
		{
			Image:       "/static/jellyfin.png",
			Title:       "Jellyfin",
			Description: "Media-Server für Filme, Serien, Bücher, Hörbücher und Fotos",
			Link:        "http://192.168.2.82:8096",
		},
		{
			Image:       "/static/plex.jpg",
			Title:       "Plex",
			Description: "Media-Server für Filme und Serien auf dem Fernseher",
			Link:        "http://192.168.2.82:32400/web",
		},
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"features": features,
	})
}
