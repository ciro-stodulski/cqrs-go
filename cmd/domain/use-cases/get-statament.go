package usecase

import (
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/entities/transaction"
)

type GetStatementUseCase interface {
	perform(accountId entity.ID) *[]transaction.Transaction
}
