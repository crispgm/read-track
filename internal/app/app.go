package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/crispgm/read-track/internal/infra"
	"github.com/crispgm/read-track/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/osteele/liquid"
	"gorm.io/gorm"
)

// Application globals
type Application struct {
	path string

	conf *infra.Conf
	db   *gorm.DB

	loc *time.Location
}

// Init globals
func (app *Application) Init(workingPath string) error {
	var err error
	app.path = workingPath
	app.conf, err = infra.LoadConf(workingPath)
	if err != nil {
		return err
	}
	app.db, err = infra.LoadDB(app.conf.DB)
	if err != nil {
		return err
	}

	app.loc, err = time.LoadLocation(app.conf.Timezone)
	if err != nil {
		return err
	}

	if !app.conf.IsDev() {
		gin.SetMode(gin.ReleaseMode)
	}
	return nil
}

// Conf getter
func (app Application) Conf() *infra.Conf {
	return app.conf
}

// DB returns DB instance
func (app Application) DB() *gorm.DB {
	if app.conf.IsDev() {
		return app.db.Debug()
	}
	return app.db
}

// AutoMigrate migrate gorm DB
func (app Application) AutoMigrate() error {
	err := app.db.AutoMigrate(
		&model.Article{},
	)
	return err
}

// CheckToken checks whether it's valid token
func (app Application) CheckToken(token string) bool {
	if token == app.conf.HTTP.AuthUser.Token {
		return true
	}

	return false
}

func (app Application) getTemplate(template string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("%s/templates/%s", app.path, template))
}

// RenderHTML with Liquid
func (app Application) RenderHTML(c *gin.Context, template string, bindings liquid.Bindings) error {
	engine := liquid.NewEngine()
	tpl, err := app.getTemplate(template)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	content, err := engine.ParseAndRender(tpl, bindings)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	if layout, ok := bindings["layout"]; ok {
		if layoutName, ok := layout.(string); ok {
			layoutTpl, err := app.getTemplate(fmt.Sprintf("layout/%s.liquid", layoutName))
			if err != nil {
				log.Fatalln(err)
				return err
			}
			bindings["content"] = content
			contentWithLayout, err := engine.ParseAndRender(layoutTpl, bindings)
			if err != nil {
				log.Fatalln(err)
				return err
			}
			c.Data(http.StatusOK, "text/html; charset=utf-8", contentWithLayout)
			return err
		}
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", content)
	return err
}
