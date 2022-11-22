package app

import (
	"github.com/gin-gonic/gin"
	"github.com/osteele/liquid"
)

// Setup .
func (app Application) Setup(c *gin.Context) {
	app.RenderHTML(c, "setup.liquid", liquid.Bindings{
		"layout":    "page",
		"path":      "/page/setup",
		"title":     "Setup",
		"pageTitle": "Setup",
		"highlight": true,

		"hostname": c.Request.Host,
		"token":    app.conf.HTTP.AuthUser.Token,
	})
}
