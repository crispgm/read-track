package app

import (
	"net/http"

	"github.com/crispgm/read-track/internal/model"
	"github.com/gin-gonic/gin"
)

// AddParams .
type AddParams struct {
	Token       string `form:"token" binding:"required"`
	Title       string `form:"title" binding:"required"`
	URL         string `form:"url" binding:"required,url"`
	Type        string `form:"type" binding:"required,oneof=read unread skip skim"`
	Author      string `form:"author"`
	Description string `form:"description"`
	Device      string `form:"device"`
}

// Add implementation of adding an article
func (app Application) Add(c *gin.Context) {
	resp := NewResponse()

	var params AddParams
	err := c.BindQuery(&params)
	if err != nil {
		c.JSON(http.StatusOK, NewError(ErrCodeParams, err).Response())
		return
	}

	article := &model.Article{
		Title:       params.Title,
		URL:         params.URL,
		Description: trimDefault(params.Description),
		Author:      trimDefault(params.Author),
		Device:      trimDefault(params.Device),
		ReadType:    params.Type,
	}
	err = model.CreateArticle(app.DB(), article)
	if err != nil {
		c.JSON(http.StatusOK, NewError(ErrCodeDBFailed, err).Response())
		return
	}
	resp.Data = article
	c.JSON(http.StatusOK, resp)
}

func trimDefault(input string) string {
	if input == "-" {
		return ""
	}
	return input
}
