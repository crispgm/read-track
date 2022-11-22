package app

import (
	"net/http"
	"time"

	"github.com/crispgm/read-track/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/osteele/liquid"
)

// ListParams .
type ListParams struct {
	Year  int `form:"year"`
	Month int `form:"month"`
}

// List .
func (app Application) List(c *gin.Context) {
	var params ListParams
	err := c.BindQuery(&params)
	if err != nil {
		c.JSON(http.StatusOK, NewError(ErrCodeParams, err).Response())
		return
	}

	if params.Year < 1970 || params.Year > 2100 {
		params.Year = time.Now().Year()
	}
	if params.Month == 0 || params.Month > 12 {
		params.Month = int(time.Now().Month())
	}

	articles, err := model.ListArticles(app.DB(), app.loc, params.Year, params.Month)
	if err != nil {
		c.JSON(http.StatusOK, NewError(ErrCodeDBFailed, err).Response())
		return
	}
	for i, a := range articles {
		articles[i].CreatedAtText = a.CreatedAt.Format("2006-01-02 15:04")
		desc := []rune(a.Description)
		if len(desc) > 140 {
			desc = desc[0:140]
			articles[i].DescriptionText = string(desc) + " ..."
		}
	}
	app.RenderHTML(c, "list.liquid", liquid.Bindings{
		"layout":    "page",
		"path":      "/page/list",
		"title":     "List",
		"pageTitle": "What I have read?",
		"articles":  articles,
	})
}
