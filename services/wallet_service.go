package services

import (
	"Labora-Wallet/models"
)

type WalletService struct {
	DbHandler models.DBHandler
}

func (w *WalletService) ProcessWalletRequest(s models.Wallet) error {

	_, err := w.DbHandler.CreateWallet(s)

	return err
}
