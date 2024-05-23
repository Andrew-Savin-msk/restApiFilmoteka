package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	Port   string `toml:"port"`
	DbPath string `toml:"db_path"`
}

// ConfigLoad loads config from toml file placed by path defined in env (CONFIG_PATH)
func (c *Config) Load() {
	configPathEnv := "CONFIG_PATH"
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load environment. Ended with error: ", err)
	}

	configPath := os.Getenv(configPathEnv)
	if configPath == "" {
		log.Fatal("Unable to load config path. Ended with error: ", err)
	}

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatal("Config file doesnt exists")
	}

	_, err = toml.DecodeFile(configPath, c)
	if err != nil {
		log.Fatal("Unable to decode file. Ended with error: ", err)
	}
}
