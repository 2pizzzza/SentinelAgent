package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/2pizzzza/sentinetAgent/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func New(cnf *config.Config) (*Postgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cnf.Database.User, cnf.Database.Password, cnf.Database.Host, cnf.Database.Port, cnf.Database.DBName))
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Postgres{
		pool: pool,
	}, nil
}

func (p *Postgres) Close() {
	p.pool.Close()
}

func (p *Postgres) Pool() *pgxpool.Pool {
	return p.pool
}
