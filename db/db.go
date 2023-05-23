package db

import (
	"database/sql"

	"Labora-Wallet/models"
)

type PostgresDBHandler struct {
	db *sql.DB
}

func (p *PostgresDBHandler) CreateWallet(w models.Wallet) error {
	// Implementar la lógica para crear un artículo en la base de datos PostgreSQL
}

func (p *PostgresDBHandler) UpdateWallet(id int) (Item, error) {
	// Implementar la lógica para obtener un artículo de la base de datos PostgreSQL
}

func (p *PostgresDBHandler) DeleteWallet(item Item) error {
	// Implementar la lógica para actualizar un artículo en la base de datos PostgreSQL
}

func (p *PostgresDBHandler) WalletStatus(id int) error {
	// Implementar la lógica para eliminar un artículo de la base de datos PostgreSQL
}
