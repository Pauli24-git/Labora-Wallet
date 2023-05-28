package models

type WalletDBHandler interface {
	CreateWallet(w Wallet) (int, error)
	UpdateWallet(id int) (Wallet, error)
	DeleteWallet(id int) error
	WalletStatus(id int) (int, error)
}

type LogsDBHandler interface {
	CreateLog(DNI int, score int, wallet_id int) error
}
