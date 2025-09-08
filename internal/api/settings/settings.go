package settings

import (
	"net/http"
	"os/exec"
	"pi-dashboard/internal/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

const (
	ErrorKeyReboot = "errorsSettingsReboot"
	ErrorKeyUmount = "errorsSettingsUmount"
)

func SettingsPage(c *gin.Context) {
	session := sessions.Default(c)

	data, err := popSessionValues(session, ErrorKeyReboot, ErrorKeyUmount)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to save session: %v", err)
		return
	}

	basePath := c.MustGet("config").(*config.Config).MntPath
	devices, err := listDirectories(basePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to read directory: %v", err)
		return
	}

	data["csrf"] = csrf.GetToken(c)
	data["devices"] = devices
	c.HTML(http.StatusOK, "settings.html", data)
}

func Reboot(c *gin.Context) {
	session := sessions.Default(c)

	cmd := exec.Command("sudo", "reboot")
	if err := cmd.Run(); err != nil {
		session.Set(ErrorKeyReboot, err.Error())
		_ = session.Save()
		c.Redirect(http.StatusSeeOther, "/settings")
		return
	}

	c.Redirect(http.StatusSeeOther, "/settings")
}

func Umount(c *gin.Context) {
	session := sessions.Default(c)

	drive := c.Query("drive")
	if drive == "" {
		session.Set(ErrorKeyUmount, "missing url query parameter: drive")
		_ = session.Save()
		c.Redirect(http.StatusSeeOther, "/settings")
		return
	}

	cmd := exec.Command("sudo", "umount", "/mnt/"+drive)
	if err := cmd.Run(); err != nil {
		session.Set(ErrorKeyUmount, err.Error())
		_ = session.Save()
		c.Redirect(http.StatusSeeOther, "/settings")
		return
	}

	c.Redirect(http.StatusSeeOther, "/settings")
}
