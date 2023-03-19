package services

import (
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/entities/account"
)

type AccountService interface {
	GetAccountById(id entity.ID) (*account.Account, error)
	GetAccountByDocument(document string) (*account.Account, error)
}
