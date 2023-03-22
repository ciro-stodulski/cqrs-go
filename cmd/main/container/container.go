package container

import (
	appcommmands "cqrs-go/cmd/application/commands"
	appsevents "cqrs-go/cmd/application/events"
	commandbillpaymenthandler "cqrs-go/cmd/application/handlers/command-bill-payment"
	eventnotificationtransactionhandler "cqrs-go/cmd/application/handlers/event-notification-transaction"
	querygetstatementhandler "cqrs-go/cmd/application/handlers/query-get-statament"
	appqueries "cqrs-go/cmd/application/queries"
	"cqrs-go/cmd/domain/bus"
	busevent "cqrs-go/cmd/main/bus-event"
	"cqrs-go/cmd/main/container/factories"
)

type (
	Container struct {
		Commands map[string]bus.CommandHandler
		Query    map[string]bus.QueryHandle
		Event    map[string]bus.EventHandle
		Bus      bus.Bus
	}
)

func New() *Container {
	infraContext := factories.MakeInfraContext()
	service := factories.MakeServiceContext(infraContext)
	evenBus := busevent.NewEventBus()

	c := &Container{
		Commands: map[string]bus.CommandHandler{
			appcommmands.CreateBillPaymentCommandName: commandbillpaymenthandler.New(service.AccountService, service.BillPaymentService, evenBus),
		},
		Query: map[string]bus.QueryHandle{
			appqueries.GetStatementQueryName: querygetstatementhandler.New(service.AccountService, service.StatementService),
		},
		Event: map[string]bus.EventHandle{
			appsevents.TransactionEventName: eventnotificationtransactionhandler.New(service.NotificationService),
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

func (c *Container) ResolveEvent(name string, kind any) bus.EventHandle {
	t := c.Event[name]

	return t
}
