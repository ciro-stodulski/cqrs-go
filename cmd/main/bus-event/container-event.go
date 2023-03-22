package busevent

import (
	appsevents "cqrs-go/cmd/application/events"
	eventnotificationtransactionhandler "cqrs-go/cmd/application/handlers/event-notification-transaction"
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/main/container/factories"
)

type containerEvent struct {
	Event map[string]bus.EventHandle
}

func newContainerEvent() *containerEvent {
	infraContext := factories.MakeInfraContext()
	service := factories.MakeServiceContext(infraContext)

	c := &containerEvent{
		Event: map[string]bus.EventHandle{
			appsevents.TransactionEventName: eventnotificationtransactionhandler.New(service.NotificationService),
		},
	}

	return c
}

func (c *containerEvent) ResolveEvent(name string, kind any) bus.EventHandle {
	t := c.Event[name]

	return t
}
