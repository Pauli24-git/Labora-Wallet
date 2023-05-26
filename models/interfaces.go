package models

type DBHandler interface {
	CreateWallet(w Wallet) (*int, error)
	UpdateWallet(id int) (Wallet, error)
	DeleteWallet(id int) error
	WalletStatus(id int) (*int, error)
}
