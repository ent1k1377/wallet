package dto

import (
	"github.com/ent1k1377/wallet/internal/database/postgres/repository"
	"github.com/shopspring/decimal"
)

type SendAmountRequest struct {
	From   string          `json:"from"`
	To     string          `json:"to"`
	Amount decimal.Decimal `json:"amount"`
}

type BalanceResponse struct {
	Balance string `json:"balance"`
}

type TransfersResponse struct {
	Transfers []repository.Transfer `json:"transfers"`
}
