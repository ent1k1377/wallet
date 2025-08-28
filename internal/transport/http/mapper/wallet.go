package mapper

import (
	"github.com/ent1k1377/wallet/internal/database/postgres/repository"
	"github.com/ent1k1377/wallet/internal/transport/http/dto"
	"github.com/shopspring/decimal"
)

func ToBalanceResponse(balance decimal.Decimal) dto.BalanceResponse {
	return dto.BalanceResponse{
		Balance: balance.String(),
	}
}

func ToTransfersResponse(transfers []repository.Transfer) dto.TransfersResponse {
	return dto.TransfersResponse{
		Transfers: transfers,
	}
}
