package postgres

import (
	"context"
	"fmt"
	"github.com/ent1k1377/wallet/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(cfg *config.DatabaseConfig) *DB {
	fmt.Println("create db")
	pool, err := pgxpool.New(context.Background(), cfg.DSN())
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	fmt.Println("connect to database")
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

func (db *DB) Close(_ context.Context) error {
	db.pool.Close()
	return nil
}
