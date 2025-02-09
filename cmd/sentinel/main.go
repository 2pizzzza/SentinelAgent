package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/2pizzzza/sentinetAgent/internal/collector/metrics"
	"github.com/2pizzzza/sentinetAgent/internal/core"
	"github.com/2pizzzza/sentinetAgent/internal/storage/postgres"
	"github.com/2pizzzza/sentinetAgent/internal/storage/redis"
	"github.com/2pizzzza/sentinetAgent/pkg/logger"

	"github.com/2pizzzza/sentinetAgent/internal/config"
)

func main() {

	cnf, err := config.New("config/config.yml")
	if err != nil {
		panic(err)
	}

	_, _ = context.WithCancel(context.Background())
	log := logger.New(cnf.Env)

	application := core.New(log, *cnf)

	_, err = redis.New(cnf.Redis.Host, cnf.Redis.Port, cnf.Redis.Password)
	if err != nil {
		panic(err)
	}

	_, err = postgres.New(cnf)
	if err != nil {
		panic(err)
	}

	linuxMetrics, err := metrics.NewLinuxMetrics()

	metricsCh := make(chan *metrics.Metrics)

	go linuxMetrics.StartCollecting(100*time.Millisecond, metricsCh)

	// go SaveMetrics(ctx, log, metricsCh, redisConn, postgresConn)

	go application.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal:", sign.String()))

	application.Stop()

	log.Info("Server is dead")
}

// func PrintMetrics(log *slog.Logger, metricsCh chan *metrics.Metrics) {
// 	for {
// 		select {
// 		case m := <-metricsCh:
// 			log.Info("Metrics", slog.Float64("CPU Usage", m.CPUUsage))
// 		}
// 	}

// }

// func SaveMetrics(ctx context.Context, log *slog.Logger, metricsCh chan *metrics.Metrics, redis *redis.Redis, postgres *postgres.Postgres) {
// 	for {
// 		select {
// 		case m := <-metricsCh:
// 			log.Info("Metrics", "CPU Usage", m.CPUUsage)

// 			data, err := json.Marshal(m)
// 			if err != nil {
// 				log.Info("Error serializing metrics:", err)
// 				continue
// 			}

// 			err = redis.Client().Publish(ctx, "metrics", data).Err()
// 			if err != nil {
// 				log.Info("Error saving metrics to Redis:", err)
// 			}
// 		}
// 	}
// }
