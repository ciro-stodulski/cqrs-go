package appcommmands

import (
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/domain/params"
	"time"
)

type createBillPaymentCommand struct {
	data      params.CreateBillPaymentParams
	timestamp string
	name      string
}

var CreateBillPaymentCommandName = "CreateBillPaymentCommand"

func NewCreateBillPaymentCommand(data params.CreateBillPaymentParams) bus.Command {
	return &createBillPaymentCommand{
		data: data, timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		name: CreateBillPaymentCommandName,
	}
}

func (c *createBillPaymentCommand) Type() any {
	return c
}

func (c *createBillPaymentCommand) Name() string {
	return c.name
}

func (c *createBillPaymentCommand) Data() any {
	return c.data
}

func (c *createBillPaymentCommand) Timestamp() string {
	return c.timestamp
}
