package app

import (
	"github.com/crispgm/read-track/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/osteele/liquid"
)

// Dashboard .
func (app Application) Dashboard(c *gin.Context) {
	errMsg := ""
	stats, err := model.GetArticleStatistics(app.DB(), app.loc)
	if err != nil {
		errMsg = err.Error()
	}
	app.RenderHTML(c, "dashboard.liquid", liquid.Bindings{
		"layout":    "page",
		"path":      "/page/dashboard",
		"title":     "Dashboard",
		"pageTitle": "Dashboard",

		"instance": app.conf.Instance,
		"timezone": app.conf.Timezone,
		"username": app.conf.HTTP.AuthUser.Name,

		"error": errMsg,
		"stats": stats,
	})
}
