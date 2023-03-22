package statementservice

import (
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/entities/transaction"
	"cqrs-go/cmd/domain/enums"
	"log"
)

func (s *statementService) GetStatementByAccountId(id entity.ID) (*[]transaction.Transaction, error) {
	log.Default().Println("Getting statement by account id: " + id.String())

	return &[]transaction.Transaction{
		{
			AccountId: id,
			Value:     100,
			Kind:      enums.Bill,
			Status:    enums.Successful,
		},
		{
			AccountId: id,
			Value:     -100,
			Kind:      enums.Pix,
			Status:    enums.Successful,
		},
	}, nil
}
