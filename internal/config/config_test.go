package config_test

import (
	"os"
	"testing"

	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/app/config"
	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	envConfigPath := "TEST_CONIG_PATH"
	err := os.Setenv(envConfigPath, "config/test_config.toml")
	assert.Nil(t, err)

	file, err := os.Create("../config/test_config.toml")
	assert.Nil(t, err)
	assert.NotNil(t, file)
	defer file.Close()

	encoder := toml.NewEncoder(file)
	assert.NotNil(t, encoder)

	inputCfg := config.Config{
		DbPath:   "postgresql://postgres:Sassassa12@localhost:5432/filmoteka?sslmode=disable",
		Port:     "8080",
		LogLevel: "debug",
	}

	err = encoder.Encode(inputCfg)
	assert.Nil(t, err)

	outputCfg := config.Load(envConfigPath)

	assert.EqualValues(t, inputCfg, outputCfg)
}
