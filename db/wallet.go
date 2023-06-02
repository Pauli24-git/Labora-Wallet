package db

import (
	"database/sql"
	"time"

	"Labora-Wallet/models"
)

type WalletDBHandler struct {
	Db *sql.DB
}

func (p *WalletDBHandler) WalletExists(dni int) (bool, error) {
	err := p.Db.QueryRow("SELECT dni FROM wallet WHERE dni = $1", dni).Scan(&dni)
	if err != nil {
		return false, err
	}
	return true, err
}

func (p *WalletDBHandler) CreateWallet(w models.Wallet) (int, error) {

	var id int
	createDate := time.Now()
	dateString := createDate.Format("2006-01-02")
	err := p.Db.QueryRow("INSERT INTO Wallet(dni, countryId, created) values ($1, $2, $3) RETURNING id", w.DNI, w.CountryId, dateString).Scan(&id)

	if err != nil {
		return id, err
	}
	return id, nil
}

func (p *WalletDBHandler) UpdateWallet(id int) (models.Wallet, error) {
	// Implementar la lógica para actualizar la wallet

	return models.Wallet{}, nil
}

func (p *WalletDBHandler) DeleteWallet(id int) error {
	res, err := p.Db.Exec("DELETE FROM wallet WHERE id = $1", id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return err
}

func (p *WalletDBHandler) WalletStatus(id int) (int, error) {
	err := p.Db.QueryRow("SELECT FROM wallet WHERE id = $1 RETURNING id", id).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
