package api

import (
	"pi-dashboard/internal/api/home"
	"pi-dashboard/internal/api/settings"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Homepage route
	r.GET("/", home.HomePage)

	// Settings route
	r.GET("/settings", settings.SettingsPage)
	r.POST("/settings/reboot", settings.Reboot)
	r.POST("/settings/umount", settings.Umount)
}
