package infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConf(t *testing.T) {
	conf, err := LoadConf("../../")
	if assert.NoError(t, err) {
		assert.Equal(t, conf.HTTP.Port, ":8080")
		assert.NotEmpty(t, conf.HTTP.Token)
		assert.Equal(t, conf.Instance, "read-track")
		assert.True(t, conf.IsDev())
		assert.False(t, conf.IsTesting())
		assert.False(t, conf.IsProduction())
	}
}
