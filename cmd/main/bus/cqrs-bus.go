package bus

import (
	"cqrs-go/cmd/domain/bus"
	"cqrs-go/cmd/main/container"
)

type CqrsBus struct {
	container *container.Container
}

func New(container *container.Container) bus.Bus {
	return &CqrsBus{container}
}

func (cqrs *CqrsBus) AsyncDispatchCommand(command bus.Command) {
	// need to implement
}

func (cqrs *CqrsBus) SyncDispatchCommand(command bus.Command) error {
	return cqrs.commandResolver(command).Perform(command)
}

func (cqrs *CqrsBus) DispatchEvent(event *bus.Event) error {
	// need to implement
	return nil
}

func (cqrs *CqrsBus) DispatchQuery(query bus.Query) (any, error) {
	return cqrs.queryResolver(query).Perform(query)
}
