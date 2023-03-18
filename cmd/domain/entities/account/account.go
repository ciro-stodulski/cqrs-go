package account

import entity "cqrs-go/cmd/domain/entities"

type Account struct {
	Id       entity.ID
	Name     string
	Balance  int
	Document string
}

func New(id entity.ID, name string, balance int, document string) *Account {
	return &Account{
		Id:       id,
		Name:     name,
		Balance:  balance,
		Document: document,
	}
}
