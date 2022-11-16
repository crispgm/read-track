package app

import (
	"sync"

	"github.com/crispgm/read-track/internal/infra"
	"github.com/crispgm/read-track/model"
	"gorm.io/gorm"
)

// Application globals
type Application struct {
	path string

	conf *infra.Conf
	db   *gorm.DB

	mu *sync.RWMutex
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
	return nil
}

// Conf getter
func (app Application) Conf() *infra.Conf {
	return app.conf
}

// MigrateDB migrate gorm DB
func (app Application) MigrateDB() error {
	err := app.db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&model.Article{},
		)
	return err
}

// IsAuthToken checks whether it's valid token
func (app Application) CheckToken(token string) bool {
	app.mu.RLock()
	defer app.mu.Unlock()
	if _, ok := app.conf.HTTP.AuthTokens[token]; ok {
		return true
	}

	return false
}
