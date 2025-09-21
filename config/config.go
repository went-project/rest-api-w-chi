package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBDialect  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

var Defaults = map[string]string{
	"PORT":        "8080",
	"DB_DIALECT":  "postgres",
	"DB_HOST":     "localhost",
	"DB_PORT":     "5432",
	"DB_USER":     "postgres",
	"DB_PASSWORD": "password",
	"DB_NAME":     "database",
	"JWT_SECRET":  "changeme",
}

func (c *Config) Initialize() {
	// Sistem ortamını al (varsayılan: local)
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	// Ortama göre .env dosyasını belirle
	var envFile string
	switch env {
	case "local":
		envFile = ".env.local"
	case "development":
		envFile = ".env.development"
	case "test":
		envFile = ".env.test"
	default: // local
		envFile = ".env"
	}

	// .env dosyasını yükle
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("Warning: Could not load %s file: %v", envFile, err)
		// Varsayılan .env dosyasını dene
		if envFile != ".env" {
			err = godotenv.Load(".env")
			if err != nil {
				log.Printf("Warning: Could not load .env file: %v", err)
			}
		}
	}

	// Konfigürasyon değerlerini yükle
	c.Port = c.Load("PORT")
	c.DBDialect = c.Load("DB_DIALECT")
	c.DBHost = c.Load("DB_HOST")
	c.DBPort = c.Load("DB_PORT")
	c.DBUser = c.Load("DB_USER")
	c.DBPassword = c.Load("DB_PASSWORD")
	c.DBName = c.Load("DB_NAME")
	c.JWTSecret = c.Load("JWT_SECRET")

	fmt.Printf("Environment: %s\n", env)
}

func (c *Config) Load(e string) string {
	var data string
	data = os.Getenv(e)
	if data == "" {
		data = Defaults[e]
	}
	return data
}
