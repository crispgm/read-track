package app

import (
	"github.com/crispgm/read-track/internal/infra"
	"github.com/gin-gonic/gin"
)

// LoadRoutes .
func (app Application) LoadRoutes(r *gin.Engine, conf *infra.Conf) {
	api := r.Group("/api")
	{
		api.GET("/add", Add)
		api.GET("/get", Get)
	}
	page := r.Group("/page", gin.BasicAuth(conf.HTTP.AuthUsers))
	{
		page.GET("/dashboard")
		page.GET("/stats")
	}
}
