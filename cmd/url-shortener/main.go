package main

import (
	"log/slog"
	"os"
	"url-shortnener/internal/config"
	"url-shortnener/internal/lib/logger/sl"
	"url-shortnener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()
	// TODO: init logger: slog
	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug message are enabled")
	// TODO init storage: sqlite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
		os.Exit(1)
	}
	url, err := storage.GetURL("google")
	if err != nil {
		log.Error("failed to fetch google", sl.Err(err))
		os.Exit(1)
	}
	log.Info("fetched url", slog.String("url", url))

	url, err = storage.GetURL("google")
	if err != nil {
		log.Error("failed to fetch google", sl.Err(err))
		os.Exit(1)
	}
	log.Info("fetched url", slog.String("url", url))

	err = storage.DeleteURL("google")
	if err != nil {
		log.Error("failed to delete google", sl.Err(err))
		os.Exit(1)
	}
	log.Info("Delete google")
	_ = storage
	// TODO: init router: chi, "chi render"

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
