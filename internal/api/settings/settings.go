package settings

import (
	"fmt"
	"net/http"
	"os/exec"
	"pi-dashboard/internal/config"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

const (
	ErrorGeneral     = "errorsGeneral"
	ErrorKeyShutdown = "errorsSettingsShutdown"
	ErrorKeyReboot   = "errorsSettingsReboot"
	ErrorKeyUmount   = "errorsSettingsUmount"
	ErrorKeyMount    = "errorsSettingsMount"
)

func Get(c *gin.Context) {
	renderPage(c, nil)
}

func Post(c *gin.Context) {
	switch c.Query("method") {
	case "shutdown":
		Shutdown(c)
	case "reboot":
		Reboot(c)
	case "umount":
		Umount(c)
	case "mount":
		Mount(c)
	default:
		renderPage(c, map[string]any{ErrorGeneral: "missing/invalid query-parameter 'method'"})
	}
}

func renderPage(c *gin.Context, data map[string]any) {
	if data == nil {
		data = make(map[string]any)
	}

	data["csrf"] = csrf.GetToken(c)

	basePath := c.MustGet("config").(*config.Config).MntPath
	devices, err := listDirectories(basePath)
	if err != nil {
		data[ErrorGeneral] = fmt.Sprintf("failed to read directory: %v", err)
		c.HTML(http.StatusOK, "settings.html", data)
		return
	}
	data["devices"] = devices

	c.HTML(http.StatusOK, "settings.html", data)
}

func Shutdown(c *gin.Context) {
	cmd := exec.Command("shutdown", "now")
	if err := cmd.Run(); err != nil {
		renderPage(c, map[string]any{ErrorKeyShutdown: err.Error()})
		return
	}

	renderPage(c, nil)
}

func Reboot(c *gin.Context) {
	cmd := exec.Command("reboot")
	if err := cmd.Run(); err != nil {
		renderPage(c, map[string]any{ErrorKeyReboot: err.Error()})
		return
	}

	renderPage(c, nil)
}

func Umount(c *gin.Context) {
	drive := c.Query("drive")
	if drive == "" {
		renderPage(c, map[string]any{ErrorKeyUmount: "missing url query parameter: drive"})
		return
	}

	cmd := exec.Command("umount", "/mnt/"+drive)
	if err := cmd.Run(); err != nil {
		renderPage(c, map[string]any{ErrorKeyUmount: err.Error()})
		return
	}

	renderPage(c, nil)
}

func Mount(c *gin.Context) {
	cmd := exec.Command("mount", "-a")
	if err := cmd.Run(); err != nil {
		renderPage(c, map[string]any{ErrorKeyMount: err.Error()})
		return
	}

	renderPage(c, nil)
	
}

