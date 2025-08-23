package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
	"time"
	"wallet/internal/pkg/math"
)

var (
	WalletNotExist = errors.New("wallet does not exist")
)

type Wallet struct {
	pool *pgxpool.Pool
}

func NewWallet(pool *pgxpool.Pool) *Wallet {
	return &Wallet{
		pool: pool,
	}
}

func (w *Wallet) Send(from, to uuid.UUID, amount decimal.Decimal) error {
	tx, err := w.pool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	query := `UPDATE wallets SET balance = balance - $1 WHERE id = $2`
	_, err = tx.Exec(context.Background(), query, amount, from.String())
	if err != nil {
		return err
	}

	query = `UPDATE wallets SET balance = balance + $1 WHERE id = $2`
	_, err = tx.Exec(context.Background(), query, amount, to.String())
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func (w *Wallet) Exist(address uuid.UUID) bool {
	var exist bool
	query := "SELECT EXISTS(SELECT 1 FROM wallets WHERE id = $1)"
	w.pool.QueryRow(context.Background(), query, address).Scan(&exist)
	return exist
}

func (w *Wallet) GetBalance(address uuid.UUID) (*decimal.Decimal, error) {
	var balance *decimal.Decimal
	query := "SELECT balance FROM wallets WHERE id = $1"
	err := w.pool.QueryRow(context.Background(), query, address).Scan(balance)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (w *Wallet) GetLastTransfers(count int) ([]Transfer, error) {
	query := "SELECT * FROM transfers ORDER BY DESC LIMIT $1"
	rows, err := w.pool.Query(context.Background(), query, count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transfers := make([]Transfer, 0, count)
	for rows.Next() {
		var t Transfer
		err := rows.Scan(&t.FromID, &t.ToID, &t.Amount, &t.CreatedAt)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transfers, nil
}

func (w *Wallet) AddRandomWallets(count int) error {
	tx, err := w.pool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	query := `INSERT INTO wallets (balance) VALUES ($1)`
	for i := 0; i < count; i++ {
		balance := math.RandomInRange(10, 1000)
		tx.Exec(context.Background(), query, balance)
	}

	tx.Commit(context.Background())
	return nil
}

func (w *Wallet) CountWallets() (int, error) {
	query := "SELECT COUNT(*) FROM wallets"
	var count int
	err := w.pool.QueryRow(context.Background(), query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

type Transfer struct {
	FromID    uuid.UUID
	ToID      uuid.UUID
	Amount    decimal.Decimal
	CreatedAt time.Time
}
