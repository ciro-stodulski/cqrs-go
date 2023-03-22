package accountservice

import (
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/entities/account"
	"log"
)

func (aService *accountService) GetAccountById(id entity.ID) (*account.Account, error) {
	log.Default().Println("Getting account by id: " + id.String())

	return nil, nil
}
