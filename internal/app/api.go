package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TrackParams .
type TrackParams struct {
	Token       string `binding:"required"`
	Title       string `binding:"required"`
	URL         string `binding:"required,url"`
	Type        string `binding:"required,oneof=read skip skim"`
	Author      string
	Description string
	Image       string
}

// Add implementation of adding an article
func Add(c *gin.Context) {
	resp := NewResponse()

	var params TrackParams
	err := c.BindQuery(&params)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	}

	// TODO: store

	c.JSON(http.StatusOK, resp)
}

// Get implementation of getting articles
func Get(c *gin.Context) {
	resp := NewResponse()
	c.JSON(http.StatusOK, resp)
}
