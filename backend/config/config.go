package config

import "os"

type Config struct {
	DBHost              string
	DBPort              string
	DBUser              string
	DBPassword          string
	DBName              string
	JWTSecret           string
	AdminPhone          string
	AdminPassword       string
	Port                string
	TelegramBotToken    string
	TelegramChatID      string
	ContactTelegramChatID string
}

func Load() *Config {
	return &Config{
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBUser:           getEnv("DB_USER", "clinic"),
		DBPassword:       getEnv("DB_PASSWORD", ""),
		DBName:           getEnv("DB_NAME", "clinic"),
		JWTSecret:        getEnv("JWT_SECRET", ""),
		AdminPhone:       getEnv("ADMIN_PHONE", ""),
		AdminPassword:    getEnv("ADMIN_PASSWORD", ""),
		Port:             getEnv("PORT", "8080"),
		TelegramBotToken:      getEnv("TELEGRAM_BOT_TOKEN", ""),
		TelegramChatID:        getEnv("TELEGRAM_CHAT_ID", ""),
		ContactTelegramChatID: getEnv("CONTACT_TELEGRAM_CHAT_ID", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
