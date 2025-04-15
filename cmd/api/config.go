package main

import (
	"log/slog"
	"pmetest/internal/services/trip"
	"sync"
)

type config struct {
	HttpPort          int    `env:"HTTP_PORT, default=4444"`
	LogLevel          string `env:"LOG_LEVEL, default=DEBUG"`
	CorsAllowedOrigin string `env:"CORS_ALLOWED_ORIGIN, default=*"`
}

type application struct {
	config      config
	logger      *slog.Logger
	wg          sync.WaitGroup
	tripService *trip.TripService
}
