package server

import (
	"path/filepath"
	"pi-dashboard/internal/api"
	"pi-dashboard/internal/config"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func configureAuth(r *gin.Engine, cfg *config.Config) {
	store := cookie.NewStore([]byte(cfg.Secret))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(csrf.Middleware(csrf.Options{
		Secret: cfg.Secret,
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/abstract/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/pages/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func Run(cfg *config.Config) error {
	r := gin.Default()

	// to prevent warning
	err := r.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	configureAuth(r, cfg)

	r.HTMLRender = loadTemplates("./templates")
	r.Static("/static", "./static")
	api.RegisterRoutes(r)

	return r.Run(cfg.IP + ":" + cfg.Port)

}
