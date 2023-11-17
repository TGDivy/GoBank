package main

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccout(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}
