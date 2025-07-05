package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	DBName     string
	JWTSecret  string
	JWTExpire  int
	ServerPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	jwtExpire, _ := strconv.Atoi(getEnv("JWT_EXPIRE_HOURS", "24"))

	return &Config{
		MongoURI:   getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DBName:     getEnv("DB_NAME", "rest_api_db"),
		JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpire:  jwtExpire,
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
