package service

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"wallet/internal/database/postgres/repository"
)

type Wallet struct {
	walletRepository *repository.Wallet
}

func NewWallet(walletRepository *repository.Wallet) *Wallet {
	return &Wallet{
		walletRepository: walletRepository,
	}
}

func (s *Wallet) Transfer(from, to uuid.UUID, amount decimal.Decimal) error {
	if !s.walletRepository.Exist(from) {
		return repository.WalletNotExist
	} else if !s.walletRepository.Exist(to) {
		return repository.WalletNotExist
	}

	s.walletRepository.Send(from, to, amount)
	return nil
}

func (s *Wallet) GetBalance(address uuid.UUID) (*decimal.Decimal, error) {
	if !s.walletRepository.Exist(address) {
		return nil, repository.WalletNotExist
	}

	return s.walletRepository.GetBalance(address)
}

func (s *Wallet) GetLastTransfers(count int) ([]repository.Transfer, error) {
	return s.walletRepository.GetLastTransfers(count)
}
