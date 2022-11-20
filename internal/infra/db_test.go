package infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDB(t *testing.T) {
	db, err := LoadDB(DBConf{
		Provider: "mysql",
		Name:     "read_track_dev.db",
	})
	if assert.NoError(t, err) {
		assert.NotNil(t, db)
	}
}
