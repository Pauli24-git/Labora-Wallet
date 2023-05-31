package services

import (
	"Labora-Wallet/models"
	"errors"
)

type WalletService struct {
	WalletDbHandler models.WalletDBHandler

	Api API

	Logs models.LogsDBHandler
}

const noDni = 0
const noScore = 0
const noId = 0

func (w *WalletService) ProcessWalletRequest(s models.Wallet) error {
	err := w.Api.initConfig()
	checkID, err := w.Api.ObtainCheckID(s.DNI, s.CountryId)
	if err != nil {
		return err
	}

	score, err := w.Api.ObtainScore(checkID)
	if err != nil {
		return err
	}
	var id int

	res, err := w.WalletDbHandler.WalletExists(s.DNI)
	if res {
		if score == 1 {
			id, err = w.WalletDbHandler.CreateWallet(s)
			if err != nil {
				return err
			}
			if id == 0 {
				return errors.New("id es 0 despues de CreateWallet")
			}
		}
	} else {
		return errors.New("El DNI ingresado ya esta asociado a una Wallet. ")
	}
	err = w.Logs.CreateLog(s.DNI, score, id)

	return err
}

func (w *WalletService) ProcessWalletDelete(id int) error {
	if err := w.WalletDbHandler.DeleteWallet(id); err != nil {
		return err
	}

	if err := w.Logs.CreateLog(noDni, noScore, noId); err != nil {
		return err
	}

	return nil
}
