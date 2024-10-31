package configs

import (
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	PublicHost string `env:"PUBLIC_HOST"`
	Port       string `env:"PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBHost     string `env:"DB_HOST"`
	DBPort     uint16 `env:"DB_PORT"`
	DBName     string `env:"DB_NAME"`
}

var Env = initConfig()

func initConfig() *Config {

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		slog.Warn("No .env file found; using environment variables set in the system")
		log.Fatal()
	} else if err != nil {
		slog.Error("Error checking .env file %s", err)
		return nil
	}
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file", "error", err)
		log.Fatal()
		return nil
	}

	return &Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "2003"),
		DBUser:     getEnv("DB_USER", "market"),
		DBPassword: getEnv("DB_PASSWORD", "linella"),
		DBPort:     getEnvAsUint("DB_PORT", 5432),
		DBName:     getEnv("DB_NAME", "dbLinella"),
		DBHost:     getEnv("DB_HOST", "postgres-db"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsUint(key string, fallback uint16) uint16 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return fallback
		}

		return uint16(i)
	}

	return fallback
}
