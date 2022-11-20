package app

import (
	"sync"
	"time"

	"github.com/crispgm/read-track/internal/infra"
	"github.com/crispgm/read-track/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Application globals
type Application struct {
	path string

	conf *infra.Conf
	db   *gorm.DB

	loc *time.Location
	mu  *sync.RWMutex
}

// Init globals
func (app *Application) Init(confPath string) error {
	var err error
	app.path = confPath
	app.conf, err = infra.LoadConf(confPath)
	if err != nil {
		return err
	}
	app.db, err = infra.LoadDB(app.conf.DB)
	if err != nil {
		return err
	}
	app.mu = &sync.RWMutex{}

	if !app.conf.IsDev() {
		gin.SetMode(gin.ReleaseMode)
	}
	app.loc, err = time.LoadLocation(app.conf.Timezone)
	if err != nil {
		return err
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
	app.mu.RLock()
	defer app.mu.Unlock()
	if _, ok := app.conf.HTTP.AuthTokens[token]; ok {
		return true
	}

	return false
}
