package usecase

import (
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/entities/transaction"
)

type BillPayment interface {
	perform(accountId entity.ID, value, receiverDocument string) *transaction.Transaction
}
