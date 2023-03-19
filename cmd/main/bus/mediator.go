package bus

import (
	"cqrs-go/cmd/domain/bus"
)

func (cqrs *CqrsBus) commandResolver(command bus.Command) bus.Handler {

	return cqrs.container.Resolve(command.Name(), command.Type())
}
