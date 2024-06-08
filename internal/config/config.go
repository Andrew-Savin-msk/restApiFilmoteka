package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	DbPath   string `toml:db_path`
	Port     string `toml:port`
	LogLevel string `toml:log_level`
}

func Load(envPath string) *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load .env with error: %s", err)
	}

	configPath := os.Getenv(envPath)
	if configPath == "" {
		log.Fatal("Enviromental variable doen't exists!")
	}
	var cfg *Config
	_, err = toml.Decode(configPath, cfg)
	if err != nil {
		log.Fatalf("Trouble with loading data from config file: %s", err)
	}
	return cfg
}
