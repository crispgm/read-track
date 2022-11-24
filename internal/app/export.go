package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/crispgm/read-track/internal/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

// ExportParams .
type ExportParams struct {
	Year   int    `form:"year"`
	Month  int    `form:"month"`
	Format string `form:"format" binding:"required,oneof=json yaml"`
}

// Export data
func (app Application) Export(c *gin.Context) {
	var params ExportParams
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

	articles, err := model.ExportArticles(app.DB(), app.loc, params.Year, params.Month)
	if err != nil {
		c.JSON(http.StatusOK, NewError(ErrCodeDBFailed, err).Response())
		return
	}
	exportArticles := convertArticles(articles)
	if params.Format == "yaml" {
		output := make(map[string][]model.ArticleExport)
		output["articles"] = exportArticles
		out, err := yaml.Marshal(output)
		if err != nil {
			c.JSON(http.StatusOK, NewError(ErrMarshalFailed, err).Response())
			return
		}
		fn := fmt.Sprintf("read-list-%d-%d.%s", params.Year, params.Month, params.Format)
		c.Writer.Header().Set("Content-type", "application/octet-stream")
		c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fn))
		c.Writer.Write(out)
	} else if params.Format == "json" {
		c.JSON(http.StatusOK, exportArticles)
	}
}

func convertArticles(articles []model.Article) []model.ArticleExport {
	var exportArticles []model.ArticleExport
	for _, a := range articles {
		article := model.ArticleExport{
			ReadType:    a.ReadType,
			Title:       a.Title,
			URL:         a.URL,
			Domain:      a.Domain,
			Author:      a.Author,
			Description: a.Description,
			Device:      a.Device,
			CreatedAt:   a.CreatedAt,
			UpdatedAt:   a.UpdatedAt,
		}
		exportArticles = append(exportArticles, article)
	}
	return exportArticles
}
