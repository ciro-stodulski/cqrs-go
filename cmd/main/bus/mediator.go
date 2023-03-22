package buscqrs

import (
	"cqrs-go/cmd/domain/bus"
)

func (cqrs *busCqrs) commandResolver(command bus.Command) bus.CommandHandler {
	return cqrs.container.ResolveCommand(command.Name(), command.Type())
}

func (cqrs *busCqrs) queryResolver(command bus.Query) bus.QueryHandle {
	return cqrs.container.ResolveQuery(command.Name(), command.Type())
}

func (cqrs *busCqrs) eventResolver(event bus.Event) bus.EventHandle {
	return cqrs.container.ResolveEvent(event.Name(), event.Type())
}
