package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (cfg Config, err error) {

	root, err := os.Getwd()
	if err != nil {
		return
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		cfg.DBHost = os.Getenv("DBHost")
		cfg.DBPort = os.Getenv("DBPort")
		cfg.DBUser = os.Getenv("DBUser")
		cfg.DBPassword = os.Getenv("DBPassword")
		cfg.DBName = os.Getenv("DBName")

		return cfg, nil
	}
	if err = envconfig.Process("", &cfg); err != nil {
		return
	}

	return
}
