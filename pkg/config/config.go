package config

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func SetDotEnv() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file.")
	} else {
		slog.Info("Loaded .env file.")
	}
}
