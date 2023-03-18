package transaction

import (
	entity "cqrs-go/cmd/domain/entities"
	enums "cqrs-go/cmd/domain/enums"
)

type Transaction struct {
	AccountId entity.ID
	Value     int
	Kind      enums.TransactionKind
}

func New(accountId entity.ID, value int, kind enums.TransactionKind) (*Transaction, error) {

	return &Transaction{
		AccountId: accountId,
		Value:     value,
		Kind:      kind,
	}, nil
}
