package usecase

import (
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/entities/transaction"
)

type PixPaymentUseCase interface {
	perform(accountId entity.ID, value int, key string) *transaction.Transaction
}
