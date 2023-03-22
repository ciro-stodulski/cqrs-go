package transaction

import (
	entity "cqrs-go/cmd/domain/entities"
	enums "cqrs-go/cmd/domain/enums"
)

type Transaction struct {
	AccountId entity.ID
	Value     int
	Kind      enums.TransactionKind
	Status    enums.TransactionStatus
}

func New(accountId entity.ID, value int, kind enums.TransactionKind, status enums.TransactionStatus) (*Transaction, error) {

	return &Transaction{
		AccountId: accountId,
		Value:     value,
		Kind:      kind,
		Status:    status,
	}, nil
}
