package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HostPort         string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresSSLMode  string
}

func NewConfig() *Config {

	cnf := Config{}

	cnf.HostPort = getOrDefault("HOST_PORT", "8080")
	cnf.PostgresHost = getOrDefault("POSTGRES_HOST", "localhost")
	cnf.PostgresPort = getOrDefault("POSTGRES_PORT", "5432")
	cnf.PostgresUser = getOrDefault("POSTGRES_USER", "postgres")
	cnf.PostgresPassword = getOrDefault("POSTGRES_PASSWORD", "postgres")
	cnf.PostgresDB = getOrDefault("POSTGRES_DB", "postgres")
	cnf.PostgresSSLMode = getOrDefault("POSTGRES_SSLMODE", "disable")

	return &cnf
}

func getOrDefault(key, defaultValue string) string {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}
