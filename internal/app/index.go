package app

import (
	"github.com/gin-gonic/gin"
	"github.com/osteele/liquid"
)

// Index .
func (app Application) Index(c *gin.Context) {
	app.RenderHTML(c, "index.liquid", liquid.Bindings{
		"layout": "page",
		"path":   "/index",
		"title":  "Home",
	})
}
