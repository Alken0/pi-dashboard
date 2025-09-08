package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	IP      string
	Port    string
	Secret  string
	MntPath string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, relying on environment variables")
	}

	ip := os.Getenv("IP")
	if ip == "" {
		panic("IP not set in environment or .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT not set in environment or .env file")
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		panic("SECRET not set in environment or .env file")
	}
	if len(secret) < 32 {
		panic("SECRET is too short")
	}

	mntPath := os.Getenv("MNT_PATH")
	if mntPath == "" {
		panic("MNT_PATH not set in environment or .env file")
	}

	return &Config{
		IP:      ip,
		Port:    port,
		Secret:  secret,
		MntPath: mntPath,
	}
}
