package models

type DBHandler interface {
	CreateWallet(item Item) error
	UpdateWallet(id int) (Item, error)
	DeleteWallet(item Item) error
	WalletStatus(id int) error
}
