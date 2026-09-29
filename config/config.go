package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	AppConfig *Config
)

type Config struct {
	AppPort         string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	JWTSecret       string
	JWTExpired      string
	JWTRefreshToken string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found.")
	}
	AppConfig = &Config{
		AppPort:         GetEnv("PORT", "3030"),
		DBHost:          GetEnv("DB_HOST", "localhost"),
		DBPort:          GetEnv("DB_PORT", "5433"),
		DBUser:          GetEnv("DB_USER", "postgres"),
		DBPassword:      GetEnv("DB_PASSWORD", "admin"),
		DBName:          GetEnv("DB_NAME", "project_management"),
		JWTSecret:       GetEnv("JWT_SECRET", "supersecret"),
		JWTRefreshToken: GetEnv("REFRESH_TOKEN_EXPIRED", "24h"),
		JWTExpired:      GetEnv("JWT_EXPIRED", "1h"),
	}
}

func GetEnv(key string, fallback string) string {
	value, exist := os.LookupEnv(key)
	if exist {
		return value
	} else {
		return fallback
	}
}

func ConnectDB() {
	cfg := AppConfig
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
}
