package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"pmetest/internal/services/trip"
	"pmetest/internal/version"
	"runtime/debug"

	"github.com/sethvargo/go-envconfig"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	var cfg config
	err := envconfig.Process(context.Background(), &cfg)
	if err != nil {
		panic(err)
	}
	showVersion := flag.Bool("version", false, "display version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	tripService := trip.New(logger)
	app := &application{
		config:      cfg,
		logger:      logger,
		tripService: tripService,
	}

	return app.serveHTTP()
}
