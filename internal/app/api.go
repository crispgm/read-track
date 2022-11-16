package app

import (
	"fmt"
	"net/http"

	"github.com/crispgm/read-track/model"
	"github.com/gin-gonic/gin"
)

// TrackParams .
type TrackParams struct {
	Token       string `form:"token" binding:"required"`
	Title       string `form:"title" binding:"required"`
	URL         string `form:"url" binding:"required,url"`
	Type        string `form:"type" binding:"required,oneof=read skip skim"`
	Author      string `form:"author"`
	Description string `form:"description"`
	Image       string `form:"image"`
}

// Add implementation of adding an article
func (app Application) Add(c *gin.Context) {
	resp := NewResponse()

	var params TrackParams
	err := c.BindQuery(&params)
	fmt.Println(err)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		return
	}

	article := &model.Article{
		Title: params.Title,
		URL:   params.URL,
		Type:  params.Type,
	}
	err = model.CreateArticle(app.db, article)
	if err != nil {
		resp.Code = -2
		resp.Message = err.Error()
		return
	}
	resp.Data = article
	c.JSON(http.StatusOK, resp)
}

// Get implementation of getting articles
func (app Application) Get(c *gin.Context) {
	resp := NewResponse()
	c.JSON(http.StatusOK, resp)
}
