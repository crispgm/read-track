package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/crispgm/read-track/internal/infra"
	"github.com/crispgm/read-track/internal/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/osteele/liquid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// Application globals
type Application struct {
	path string
	conf *infra.Conf
	db   *gorm.DB
	loc  *time.Location

	authenticator *infra.Authenticator

	logger zerolog.Logger
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

	app.authenticator, err = infra.InitAuthenticator(app.Conf().Auth0)
	if err != nil {
		return err
	}

	if !app.conf.IsDev() {
		app.logger = zerolog.New(os.Stderr).Level(zerolog.InfoLevel).With().Timestamp().Logger()
		gin.SetMode(gin.ReleaseMode)
	} else {
		app.logger = zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Logger()
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

// GetUserInfo .
func (app Application) GetUserInfo(ctx *gin.Context) *model.User {
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	if profile != nil {
		if info, ok := profile.(map[string]interface{}); ok {
			user := &model.User{
				Name:     info["name"].(string),
				Nickname: info["nickname"].(string),
				Picture:  info["picture"].(string),
				Sub:      info["sub"].(string),
			}
			return user
		}
	}
	return nil
}

// IsAuthenticated .
func (app Application) IsAuthenticated(ctx *gin.Context) {
	user := app.GetUserInfo(ctx)
	if user == nil || user.Sub != app.Conf().Auth0.UserID {
		ctx.Redirect(http.StatusSeeOther, "/login")
	} else {
		ctx.Next()
	}
}
