package services

import (
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/entities/transaction"
)

type StatementService interface {
	GetStatementByAccountId(id entity.ID) (*[]transaction.Transaction, error)
}
