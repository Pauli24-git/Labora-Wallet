package services

import "Labora-Wallet/models"

type TransactionService struct {
	TransactionDBHandler models.TransactionsDBHandler
}

func (t *TransactionService) CreateTransaction(tr models.Transactions) error {
	err := t.TransactionDBHandler.CheckIDBeforeTransaction(tr)
	if err != nil {
		return err
	}

	err = t.TransactionDBHandler.CheckFundsBeforeTransaction(tr)
	if err != nil {
		return err
	}

	err = t.TransactionDBHandler.InsertTransaction(tr)
	if err != nil {
		return err
	}

	return err
}

func (t *TransactionService) TransactionsStatusbyID(id int) (models.TransactionInfo, error) {
	res, err := t.TransactionDBHandler.TransactionsStatusbyID(id)
	if err != nil {
		return res, err
	}
	return res, err
}
