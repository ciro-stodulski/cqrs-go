package transaction

import (
	entity "cqrs-go/cmd/domain/entities"
	enums "cqrs-go/cmd/domain/enums"
)

type Transaction struct {
	AccountId entity.ID
	Value     int
	Type      enums.Transaction
}
