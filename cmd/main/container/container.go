package container

import (
	appcommands "cqrs-go/cmd/application/commands"
	billpaymenthandlers "cqrs-go/cmd/application/handlers/bill-payment-command"
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/main/container/factories"
)

type (
	Container struct {
		Bd map[string]bus.Handler
	}
)

func New() *Container {
	infraContext := factories.MakeInfraContext()

	accountService := factories.MakeServiceContext(infraContext).AccountService

	billPaymentService := factories.MakeServiceContext(infraContext).BillPaymentService

	c := &Container{
		Bd: map[string]bus.Handler{
			appcommands.CommandName: billpaymenthandlers.New(accountService, billPaymentService),
		},
	}

	return c
}

func (c *Container) Resolve(name string, kind any) bus.Handler {
	t := c.Bd[name]

	return t
}
