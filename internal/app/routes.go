package app

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/crispgm/read-track/internal/infra"
	"github.com/gin-gonic/gin"
)

// LoadRoutes .
func (app Application) LoadRoutes(r *gin.Engine, conf *infra.Conf, resources *embed.FS) {
	var tpl *template.Template
	var fs http.FileSystem
	if conf.IsDev() {
		r.LoadHTMLGlob("templates/*.tmpl")
		fs = http.Dir("static")
		r.StaticFS("/public/static", fs)
	} else {
		tpl = template.Must(template.New("").ParseFS(resources, "templates/*.tmpl"))
		r.SetHTMLTemplate(tpl)
		fs = http.FS(resources)
		r.StaticFS("/public", fs)
	}

	r.GET("/", app.Index)

	api := r.Group("/api")
	{
		api.GET("/add", app.Add)
	}
	page := r.Group("/page", gin.BasicAuth(conf.HTTP.AuthUsers))
	{
		page.GET("/export", app.Export)
		page.GET("/dashboard", app.Dashboard)
	}
}

func (app Application) validTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
