package mapper

import (
	"github.com/shopspring/decimal"
	"wallet/internal/database/postgres/repository"
	"wallet/internal/transport/http/dto"
)

func ToBalanceResponse(balance *decimal.Decimal) dto.BalanceResponse {
	return dto.BalanceResponse{
		Balance: balance.String(),
	}
}

func ToTransfersResponse(transfers []repository.Transfer) dto.TransfersResponse {
	return dto.TransfersResponse{
		Transfers: transfers,
	}
}
