package params

import (
	"cqrs-go/cmd/domain/entities/transaction"
	"time"
)

type NotificationTransactionParams struct {
	Date       time.Time
	Trasaction transaction.Transaction
}
