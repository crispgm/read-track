package app

import (
	"github.com/crispgm/read-track/internal/infra"
	"gorm.io/gorm"
)

// Application globals
type Application struct {
	path string

	conf *infra.Conf
	db   *gorm.DB
}

// Init globals
func (app *Application) Init(confPath string) error {
	var err error
	app.path = confPath
	app.conf, err = infra.LoadConf("./")
	if err != nil {
		return err
	}
	app.db, err = infra.LoadDB(app.conf.DB)
	if err != nil {
		return err
	}

	return nil
}

// Conf getter
func (app Application) Conf() infra.Conf {
	return *app.conf
}
