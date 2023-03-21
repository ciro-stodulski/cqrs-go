package bus

import (
	"cqrs-go/cmd/domain/bus"
)

func (cqrs *CqrsBus) commandResolver(command bus.Command) bus.CommandHandler {
	return cqrs.container.ResolveCommand(command.Name(), command.Type())
}

func (cqrs *CqrsBus) queryResolver(command bus.Query) bus.QueryHandle {
	return cqrs.container.ResolveQuery(command.Name(), command.Type())
}
