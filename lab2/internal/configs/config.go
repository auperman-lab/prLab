package configs

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Config struct {
	PublicHost string `env:"PUBLIC_HOST"`
	Port       string `env:"PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBAddress  string `env:"DB_ADDRESS"`
	DBName     string `env:"DB_NAME"`
}

var Env = initConfig()

func initConfig() *Config {

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		slog.Warn("No .env file found; using environment variables set in the system")
	} else if err != nil {
		slog.Error("Error checking .env file %s", err)
		return nil
	}
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file", "error", err)
		return nil
	}

	return &Config{
		PublicHost: os.Getenv("PUBLIC_HOST"),
		Port:       os.Getenv("PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBAddress:  os.Getenv("DB_ADDRESS"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
