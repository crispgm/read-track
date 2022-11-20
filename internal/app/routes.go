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
	tpl := template.Must(template.New("").ParseFS(resources, "templates/index.tmpl"))
	r.SetHTMLTemplate(tpl)
	// example: /public/assets/images/example.png
	r.StaticFS("/public", http.FS(resources))

	r.GET("/", app.Index)

	api := r.Group("/api")
	{
		api.GET("/add", app.Add)
	}
	page := r.Group("/page", gin.BasicAuth(conf.HTTP.AuthUsers))
	{
		page.GET("/export", app.Export)
		// page.GET("/dashboard")
	}
}

func (app Application) validTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
