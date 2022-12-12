package app

import (
	"github.com/crispgm/read-track/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/osteele/liquid"
)

// Dashboard .
func (app Application) Dashboard(c *gin.Context) {
	errMsg := ""
	stats, ranks, err := model.GetArticleStatistics(app.DB(), app.loc)
	if err != nil {
		errMsg = err.Error()
	}
	user := app.GetUserInfo(c)
	app.RenderHTML(c, "dashboard.liquid", liquid.Bindings{
		"layout": "page",
		"path":   "/page/dashboard",
		"title":  "Dashboard",

		"instance": app.conf.Instance,
		"timezone": app.conf.Timezone,
		"database": app.conf.DB.Name,
		"username": user.Name,
		"nickname": user.Nickname,
		"picture":  user.Picture,
		"token":    app.conf.HTTP.Token,

		"error": errMsg,
		"stats": stats,
		"ranks": ranks,
	})
}
