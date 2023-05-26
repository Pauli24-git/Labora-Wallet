package services

import (
	"Labora-Wallet/db"
	"Labora-Wallet/models"
)

type WalletService struct {
	DbHandler models.DBHandler

	Api API

	Logs db.Logs
}

func (w *WalletService) ProcessWalletRequest(s models.Wallet) error {
	checkID, err := w.Api.ObtainCheckID(s.DNI, s.CountryId)
	if err != nil {
		return err
	}

	score, err := w.Api.ObtainScore(checkID)
	if err != nil {
		return err
	}
	var id *int

	if score == 1 {
		id, err = w.DbHandler.CreateWallet(s)
		if err != nil {
			return err
		}
	}

	err = w.Logs.CreateLog(s.DNI, score, *id)

	return err
}

func (w *WalletService) ProcessWalletDelete(id int) error {
	err := w.DbHandler.DeleteWallet(id)
	if err != nil {
		return err
	}

	err = w.Logs.CreateLog(0, 0, 0)
	if err != nil {
		return err
	}

	return nil
}
