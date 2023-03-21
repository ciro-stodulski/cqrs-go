package container

import (
	appcommmands "cqrs-go/cmd/application/commands"
	billpaymentcommandhandlers "cqrs-go/cmd/application/handlers/bill-payment-command"
	getstatementqueryhandlers "cqrs-go/cmd/application/handlers/get-statament-query"
	appqueries "cqrs-go/cmd/application/queries"
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/main/container/factories"
)

type (
	Container struct {
		Commands map[string]bus.CommandHandler
		Query    map[string]bus.QueryHandle
	}
)

func New() *Container {
	infraContext := factories.MakeInfraContext()

	accountService := factories.MakeServiceContext(infraContext).AccountService

	billPaymentService := factories.MakeServiceContext(infraContext).BillPaymentService

	statementService := factories.MakeServiceContext(infraContext).StatementService

	c := &Container{
		Commands: map[string]bus.CommandHandler{
			appcommmands.CreateBillPaymentCommandName: billpaymentcommandhandlers.New(accountService, billPaymentService),
		},
		Query: map[string]bus.QueryHandle{
			appqueries.GetStatementQueryName: getstatementqueryhandlers.New(accountService, statementService),
		},
	}

	return c
}

func (c *Container) ResolveCommand(name string, kind any) bus.CommandHandler {
	t := c.Commands[name]

	return t
}

func (c *Container) ResolveQuery(name string, kind any) bus.QueryHandle {
	t := c.Query[name]

	return t
}
