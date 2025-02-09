package core

import (
	"fmt"
	"github.com/2pizzzza/sentinetAgent/internal/config"
	"github.com/2pizzzza/sentinetAgent/pkg/logger/sl"
	"log/slog"
	"net/http"
)

type App struct {
	log        *slog.Logger
	httpServer *http.Server
	host       string
	port       int
}

func New(log *slog.Logger, cnf config.Config) *App {
	_ = http.NewServeMux()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cnf.Server.Host, cnf.Server.Port),
		Handler: nil,
	}

	return &App{
		log:        log,
		httpServer: httpServer,
		host:       cnf.Server.Host,
		port:       cnf.Server.Port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "httpapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.String("port", a.host),
		slog.Int("port", a.port),
	)

	log.Info("starting server")

	if err := a.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "httpapp.Stop"

	a.log.With(slog.String("op", op)).Info("Stopping HTTP server", slog.Int("port", a.port))

	if err := a.httpServer.Close(); err != nil {
		a.log.Error("error while closing HTTP server", sl.Err(err))
	}
}
