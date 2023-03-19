package accountservice

import (
	"cqrs-go/cmd/domain/services"
)

type accountService struct {
}

func New() services.AccountService {
	return &accountService{}
}
