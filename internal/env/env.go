package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetString(key, fallback string) string {
    if err := godotenv.Load("D:\\Projects2026\\krant\\krant-backend\\.env"); err != nil {
        log.Println("No .env file found, using system env")
    }
    if val := os.Getenv(key); val != "" {
        return val
    }
    return fallback
}
