package config

import "os"

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	JWTSecret     string
	AdminPhone    string
	AdminPassword string
	Port          string
}

func Load() *Config {
	return &Config{
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "clinic"),
		DBPassword:    getEnv("DB_PASSWORD", "clinic_secret_2024"),
		DBName:        getEnv("DB_NAME", "clinic"),
		JWTSecret:     getEnv("JWT_SECRET", "clinic-jwt-secret"),
		AdminPhone:    getEnv("ADMIN_PHONE", "998901475130"),
		AdminPassword: getEnv("ADMIN_PASSWORD", "12345678"),
		Port:          getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
