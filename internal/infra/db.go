package infra

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// LoadDB load DB instance
func LoadDB(dbconf DBConf) (*gorm.DB, error) {
	var (
		dsn string
		db  *gorm.DB
		err error
	)

	if dbconf.Provider == "sqlite" {
		dsn = dbconf.Name
	}
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	return db, err
}
