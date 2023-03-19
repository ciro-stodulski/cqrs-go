package appcommands

import (
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/domain/params"
	"time"
)

type CreateBillPaymentCommand struct {
	data      params.CreateBillPaymentParams
	timestamp string
	name      string
}

var CommandName = "CreateBillPaymentCommand"

func NewCreateBillPaymentCommand(data params.CreateBillPaymentParams) bus.Command {
	return &CreateBillPaymentCommand{data: data, timestamp: time.Now().UTC().Format(time.RFC3339Nano), name: CommandName}
}

func (c *CreateBillPaymentCommand) Type() interface{} {
	return c
}

func (c *CreateBillPaymentCommand) Name() string {
	return c.name
}

func (c *CreateBillPaymentCommand) Data() interface{} {
	return c.data
}

func (c *CreateBillPaymentCommand) Timestamp() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}
