package db

import (
	"database/sql"
	"time"

	"Labora-Wallet/models"
)

type PostgresDBHandler struct {
	Db *sql.DB
}

func (p *PostgresDBHandler) CreateWallet(w models.Wallet) (*int, error) {

	var id int
	createDate := time.Now()
	dateString := createDate.Format("2006-01-02")
	err := p.Db.QueryRow("INSERT INTO Wallet(dni, countryId, Date) values ($1, $2, $3) RETURNING id", w.DNI, w.CountryId, dateString).Scan(&id)

	if err != nil {
		return &id, err
	}
	return &id, nil
	//aca es lo mismo si retorno err(vacio) o si esta bien devolver el nil cargado a mano ?
}

func (p *PostgresDBHandler) UpdateWallet(id int) (models.Wallet, error) {
	// Implementar la lógica para actualizar la wallet

	return models.Wallet{}, nil
}

func (p *PostgresDBHandler) DeleteWallet(id int) error {
	// Implementar la lógica para eliminar la wallet
	return nil
}

func (p *PostgresDBHandler) WalletStatus(id int) error {
	// Implementar la lógica para chequear el estado de la wallet
	return nil
}
