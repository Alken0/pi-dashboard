package settings

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

type Feature struct {
	Image       string
	Title       string
	Description string
	Link        string
}

func SettingsPage(c *gin.Context) {
	token := csrf.GetToken(c)

	// List directories under /mnt
	entries, err := os.ReadDir("C:/Users/jonas/Documents/pi-dashboard/cmd")
	if err != nil {
		println("err " + err.Error())
		return
	}

	// Filter only directories
	devices := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			devices = append(devices, entry.Name())
		}
	}

	c.HTML(http.StatusOK, "settings.html", gin.H{
		"csrf":    token,
		"Devices": devices,
	})
}

func Reboot(c *gin.Context) {
	session := sessions.Default(c)

	cmd := exec.Command("sudo", "reboot")

	if err := cmd.Run(); err != nil {
		session.Set("error_msg", err.Error())
		session.Save()
		c.Redirect(http.StatusSeeOther, "/error")
		return
	}

	token := csrf.GetToken(c)
	c.HTML(http.StatusOK, "settings.html", gin.H{
		"csrf": token,
	})
}

func ButtonClick2(c *gin.Context) {
	fmt.Println("Button2 clicked!")
	c.JSON(http.StatusOK, gin.H{"message": "Button2 clicked, check the terminal!"})
}
