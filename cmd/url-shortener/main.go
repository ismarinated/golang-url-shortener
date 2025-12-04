package main

import (
	"log"
	"log/slog"
	"os"
	"url-shortener/cmd/internal/config"
	"url-shortener/cmd/internal/storage/postgres"

	"github.com/joho/godotenv"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env file does not exist")
	}

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	log.Info("starting database connection")
	storage, err := postgres.New()
	if err != nil {
		log.Error("failed to init storage", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer func() {
		storage.Close()
		log.Info("database closed")
	}()
	log.Info("database connected")

	//
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return log
}
