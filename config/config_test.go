package config_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/victorcel/crud-rest-vozy/config"
	"testing"
)

func TestGetConfig(t *testing.T) {
	conf := config.GetConfig()
	assert.NotEmpty(t, conf)
}
