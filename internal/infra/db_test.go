package infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDB(t *testing.T) {
	db, err := LoadDB(DBConf{
		Provider: "mysql",
		Host:     "127.0.0.1:3306",
		Name:     "read_track_dev",
		User:     "root",
		Pass:     "",
	})
	if assert.NoError(t, err) {
		assert.NotNil(t, db)
	}
}
