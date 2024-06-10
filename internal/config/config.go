package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	Port       string `toml:"port"`
	DbType     string `toml:"db_type"`
	DbPath     string `toml:"db_path"`
	SchemaPath string `toml:"schema_path"`
	LogLevel   string `toml:"log_level"`
}

func Load(envPath string) *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load .env with error: %s", err)
	}

	configPath := os.Getenv(envPath)
	if configPath == "" {
		log.Fatal("Enviromental variable doen't exists!")
	}

	_, err = os.Stat(configPath)
	if err != nil {
		log.Fatalf("Error with file stats: %s", err)
	}

	var cfg Config
	_, err = toml.DecodeFile(configPath, &cfg)
	if err != nil {
		log.Fatalf("Trouble with loading data from config file: %s", err)
	}
	return &cfg
}
