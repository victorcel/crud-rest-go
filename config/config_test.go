package config_test

import (
	"crud-rest-vozy/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfig(t *testing.T) {
	conf := config.GetConfig()
	assert.NotEmpty(t, conf)
}
