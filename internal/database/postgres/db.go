package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"wallet/internal/config"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(cfg *config.DatabaseConfig) *DB {
	pool, err := pgxpool.New(context.Background(), cfg.DSN())
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	if err := pool.Ping(context.Background()); err != nil {
		panic("failed to ping database: " + err.Error())
	}

	return &DB{
		pool: pool,
	}
}

func (db *DB) GetPool() *pgxpool.Pool {
	return db.pool
}

func (db *DB) Close(ctx context.Context) error {
	db.pool.Close()
	return nil
}
