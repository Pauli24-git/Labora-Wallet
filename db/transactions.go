package db

import (
	"Labora-Wallet/models"
	"database/sql"
	"errors"
)

type TransactionDBHandler struct {
	Db *sql.DB
}

func (t *TransactionDBHandler) CheckIDBeforeTransaction(tr models.Transactions) error {
	count := 0
	rows, err := t.Db.Query("SELECT * FROM wallet WHERE id IN ($1, $2)", tr.SenderId, tr.ReceiverId)
	if err != nil {
		return errors.New("No se encontró el ID buscado.")
	}

	for rows.Next() {
		count++
	}

	if count != 2 {
		return errors.New("No existe alguna de las wallets")
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return err
}

func (t *TransactionDBHandler) CheckFundsBeforeTransaction(c models.Transactions) error {
	var balance int

	err := t.Db.QueryRow("SELECT balance FROM wallet WHERE id = $1", c.SenderId).Scan(&balance)
	if err != nil || balance < c.Amount {
		return errors.New("La billetera no cuenta con fondos suficientes para realizar la transaccion. ")
	}
	return nil
}

func (t *TransactionDBHandler) InsertTransaction(c models.Transactions) error {
	var id int

	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.QueryRow("INSERT INTO transactions(sender_id, receiver_id, operation, amount) values ($1, $2, $3, $4) RETURNING id", c.SenderId,
		c.ReceiverId, c.Operation, c.Amount).Scan(&id)

	if err != nil {
		return errors.New("No se pudo realizar la transaccion.")
	}

	_, err = tx.Exec("UPDATE wallet SET balance = balance + $1 WHERE id = $2", c.Amount, c.ReceiverId)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE wallet SET balance = balance - $1 WHERE id = $2", c.Amount, c.SenderId)
	if err != nil {
		return err
	}
	return err
}

func (t *TransactionDBHandler) TransactionsStatusbyID(id int) (models.TransactionInfo, error) {
	var newTransactionInfo models.TransactionInfo
	res, err := t.Db.Query("SELECT t.sender_Id, t.amount, t.created, w.balance "+
		"FROM transactions t "+
		"JOIN wallet w ON t.sender_Id = w.id OR t.receiver_Id = w.id "+
		"WHERE w.id = $1 "+
		"ORDER BY created DESC", id)
	if err != nil {
		return newTransactionInfo, errors.New("No se pudo encontrar el ID para obtener el STATUS.")
	}

	newTransactionInfo.ID = id
	for res.Next() {
		var senderid int
		var mov models.Movements
		err := res.Scan(&senderid, &mov.Amount, &mov.Time, &newTransactionInfo.Balance)
		if err != nil {
			return newTransactionInfo, err
		}
		if senderid == id {
			mov.Operation = "withdrawal"
		} else {
			mov.Operation = "deposit"
		}

		newTransactionInfo.Movements = append(newTransactionInfo.Movements, mov)
	}
	return newTransactionInfo, err
}
