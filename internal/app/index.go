package app

import (
	"github.com/gin-gonic/gin"
	"github.com/osteele/liquid"
)

// Index .
func (app Application) Index(c *gin.Context) {
	app.RenderHTML(c, "index.liquid", liquid.Bindings{
		"layout":    "page",
		"path":      "/index",
		"title":     "Home",
		"pageTitle": "Welcome to Read Track",

		"instance": app.conf.Instance,
		"timezone": app.conf.Timezone,
		"username": app.conf.HTTP.AuthUser.Name,
	})
}
