package api

import (
	"net/http"
	"pi-dashboard/internal/api/home"
	"pi-dashboard/internal/api/settings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Homepage route
	r.GET("/", home.HomePage)

	// Settings route
	r.GET("/settings", settings.SettingsPage)
	r.POST("/settings/reboot", settings.Reboot)
	r.POST("/settings/click2", settings.ButtonClick2)

	// error handling
	r.GET("/error", func(c *gin.Context) {
		session := sessions.Default(c)
		errMsg := session.Get("error_msg")
		session.Delete("error_msg")
		session.Save()

		c.HTML(http.StatusInternalServerError, "error.page.html", gin.H{
			"Error": errMsg,
		})
	})
}
