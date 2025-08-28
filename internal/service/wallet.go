package service

import (
	"github.com/ent1k1377/wallet/internal/database/postgres/repository"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	walletRepository *repository.Wallet
}

func NewWallet(walletRepository *repository.Wallet) *Wallet {
	return &Wallet{
		walletRepository: walletRepository,
	}
}

func (s *Wallet) InitializeFirstRun() error {
	cnt, err := s.walletRepository.CountWallets()
	if err != nil {
		return err
	}

	if cnt == 0 {
		err = s.walletRepository.AddRandomWallets(10)
	}

	return err
}

func (s *Wallet) Transfer(from, to uuid.UUID, amount decimal.Decimal) error {
	if !s.walletRepository.Exist(from) {
		return repository.WalletNotExist
	} else if !s.walletRepository.Exist(to) {
		return repository.WalletNotExist
	}

	return s.walletRepository.Send(from, to, amount)
}

func (s *Wallet) GetBalance(address uuid.UUID) (decimal.Decimal, error) {
	if !s.walletRepository.Exist(address) {
		return decimal.Decimal{}, repository.WalletNotExist
	}

	return s.walletRepository.GetBalance(address)
}

func (s *Wallet) GetLastTransfers(count int) ([]repository.Transfer, error) {
	return s.walletRepository.GetLastTransfers(count)
}
