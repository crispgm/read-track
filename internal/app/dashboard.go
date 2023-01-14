package app

import (
	"time"

	"github.com/crispgm/read-track/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/osteele/liquid"
)

// DashboardParams .
type DashboardParams struct {
	Year  int `form:"year"`
	Month int `form:"month"`
	Day   int `form:"day"`
}

// Dashboard .
func (app Application) Dashboard(c *gin.Context) {
	var params DashboardParams
	_ = c.BindQuery(&params)
	now := time.Now().In(app.loc)
	if params.Year < 1970 {
		params.Year = now.Year()
	}
	if params.Month == 0 || params.Month > 12 {
		params.Month = int(now.Month())
	}
	if params.Day == 0 || params.Day > 31 {
		params.Day = int(now.Day())
	}

	date := time.Date(params.Year, time.Month(params.Month), params.Day, 0, 0, 0, 0, app.loc)
	stats, ranks, err := model.GetArticleStatistics(app.DB(), app.loc, date)
	errMsg := ""
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
