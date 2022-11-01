package infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConf(t *testing.T) {
	conf, err := LoadConf("../../")
	if assert.NoError(t, err) {
		assert.Equal(t, conf.HTTP.Port, ":8080")
	}
}
