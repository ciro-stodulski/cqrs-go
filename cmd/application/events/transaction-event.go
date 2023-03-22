package appsevents

import (
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/domain/params"
	"time"
)

type transactionEvent struct {
	data      params.NotificationTransactionParams
	timestamp string
	name      string
}

var TransactionEventName = "TransactionEvent"

func NewTransactionEvent(data params.NotificationTransactionParams) bus.Event {
	return &transactionEvent{data: data, timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		name: TransactionEventName,
	}
}

func (c *transactionEvent) Type() any {
	return c
}

func (c *transactionEvent) Name() string {
	return c.name
}

func (c *transactionEvent) Data() any {
	return c.data
}

func (c *transactionEvent) Timestamp() string {
	return c.timestamp
}
