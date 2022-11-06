package infra

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// LoadDB load DB instance
func LoadDB(dbconf DBConf) (*gorm.DB, error) {
	var (
		dsn string
		db  *gorm.DB
		err error
	)

	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbconf.User,
		dbconf.Pass,
		dbconf.Host,
		dbconf.Name,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
