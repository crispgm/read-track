package app

import (
	"net/http"

	"github.com/crispgm/read-track/internal/model"
	"github.com/gin-gonic/gin"
)

// Dashboard .
func (app Application) Dashboard(c *gin.Context) {
	errMsg := ""
	stats, err := model.GetArticleStatistics(app.DB(), app.loc)
	if err != nil {
		errMsg = err.Error()
	}
	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
		"error": errMsg,
		"data":  stats,
	})
}
