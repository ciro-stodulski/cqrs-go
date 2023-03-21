package statementservice

import (
	entity "cqrs-go/cmd/domain/entities"
	"cqrs-go/cmd/domain/entities/transaction"
	"cqrs-go/cmd/domain/enums"
	"fmt"
)

func (s *statementService) GetStatementByAccountId(id entity.ID) (*[]transaction.Transaction, error) {

	fmt.Println(id)

	return &[]transaction.Transaction{
		{
			AccountId: id,
			Value:     100,
			Kind:      enums.Bill,
		},
		{
			AccountId: id,
			Value:     -100,
			Kind:      enums.Pix,
		},
	}, nil
}
