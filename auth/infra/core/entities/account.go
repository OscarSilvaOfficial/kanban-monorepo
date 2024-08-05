package entities

import "fmt"

type Account struct {
	name string
	email string
	password string
}

func (account *Account) createAccount(name, email, password string) {
	account.name = name
	account.email = email
	account.password = password
}

func CreateAccount(name, email, password string) *Account {
	account := new(Account)
	account.createAccount(name, email, password)

	fmt.Println(account.name, account.email, account.password)
	return account
}