package params

import (
	entity "cqrs-go/cmd/domain/entities"
)

type CreateBillPaymentParams struct {
	AccountId              entity.ID
	ReceiveDocumentAccount string
	Value                  int
}
