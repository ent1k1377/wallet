package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
	"wallet/internal/service"
	"wallet/internal/transport/http/dto"
	"wallet/internal/transport/http/mapper"
)

type Wallet struct {
	walletService *service.Wallet
}

func NewWallet(walletService *service.Wallet) *Wallet {
	return &Wallet{
		walletService: walletService,
	}
}

func (w *Wallet) Send(ctx *gin.Context) {
	var request dto.SendAmountRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, mapper.ToErrorResponse("json body is not valid"))
		return
	}

	uuidFrom, err := uuid.Parse(request.From)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, mapper.ToErrorResponse("uuid `from` is not valid"))
		return
	}

	uuidTo, err := uuid.Parse(request.To)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, mapper.ToErrorResponse("uuid `to` is not valid"))
		return
	}

	amount, err := decimal.NewFromString(request.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, mapper.ToErrorResponse("amount is not valid"))
		return
	}

	err = w.walletService.Transfer(uuidFrom, uuidTo, amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, mapper.ToErrorResponse("transfer error"))
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToSuccessResponse("successful money transfer"))
}

func (w *Wallet) GetLast(ctx *gin.Context) {
	countQuery := ctx.Query("count")
	count, err := strconv.Atoi(countQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, mapper.ToErrorResponse("count must be a number"))
		return
	}

	transfers, err := w.walletService.GetLastTransfers(count)
	if transfers != nil {
		ctx.JSON(http.StatusInternalServerError, mapper.ToErrorResponse("get last transfers"))
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToTransfersResponse(transfers))
}

func (w *Wallet) GetBalance(ctx *gin.Context) {
	addressParam := ctx.Param("address")
	address, err := uuid.Parse(addressParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, mapper.ToErrorResponse("address is not valid"))
		return
	}

	balance, err := w.walletService.GetBalance(address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, mapper.ToErrorResponse("get balance error"))
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToBalanceResponse(balance))
}
