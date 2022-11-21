package app

import (
	"net/http"

	"github.com/crispgm/read-track/internal/infra"
	"github.com/gin-gonic/gin"
)

// LoadRoutes .
func (app Application) LoadRoutes(r *gin.Engine, conf *infra.Conf) {
	if conf.IsDev() {
		r.StaticFS("/public/static", http.Dir("static"))
	} else {
		r.StaticFS("/public", http.FS(app.fs))
	}

	r.GET("/", app.Index)

	api := r.Group("/api")
	{
		api.GET("/add", app.Add)
	}
	page := r.Group("/page", gin.BasicAuth(conf.AuthAccounts()))
	{
		page.GET("/export", app.Export)
		page.GET("/setup", app.Setup)
		page.GET("/dashboard", app.Dashboard)
	}
}

func (app Application) validTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
