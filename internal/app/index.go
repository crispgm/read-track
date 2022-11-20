package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app Application) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
