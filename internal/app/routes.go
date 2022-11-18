package app

import (
	"github.com/crispgm/read-track/internal/infra"
	"github.com/gin-gonic/gin"
)

func (app Application) validTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// LoadRoutes .
func (app Application) LoadRoutes(r *gin.Engine, conf *infra.Conf) {
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
