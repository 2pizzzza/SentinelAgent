package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/2pizzzza/sentinetAgent/internal/collector/metrics"
	"github.com/2pizzzza/sentinetAgent/internal/core"
	"github.com/2pizzzza/sentinetAgent/pkg/logger"

	"github.com/2pizzzza/sentinetAgent/internal/config"
)

func main() {

	cnf, err := config.New("config/config.yml")
	if err != nil {
		panic(err)
	}

	log := logger.New(cnf.Env)

	application := core.New(log, *cnf)

	linuxMetrics, err := metrics.NewLinuxMetrics()
	if err != nil {
		panic(err)
	}

	go linuxMetrics.StartCollecting(100 * time.Millisecond)

	go application.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal:", sign.String()))

	application.Stop()

	log.Info("Server is dead")
}
