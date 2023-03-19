package accountservice

import (
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/entities/account"
)

func (aService *accountService) GetAccountById(id entity.ID) (*account.Account, error) {

	return nil, nil
}
