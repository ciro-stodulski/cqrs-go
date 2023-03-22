package buscqrs

import (
	"cqrs-go/cmd/domain/bus"
	eventemitter "cqrs-go/cmd/infra/adapters/event-emitter"
	"cqrs-go/cmd/main/container"
	"log"
)

type busCqrs struct {
	container *container.Container
	event     *eventemitter.EventEmitter
}

var DISPATCH_COMMAND = "DISPATCH_COMMAND"

func NewbusCqrs(container *container.Container) bus.Bus {

	busCqrs := &busCqrs{container: container, event: eventemitter.NewEventEmitter()}

	busCqrs.event.On(DISPATCH_COMMAND, func(args ...any) {
		command := args[0].(bus.Command)
		busCqrs.SyncDispatchCommand(command)
	})

	return busCqrs
}

func (cqrs *busCqrs) SyncDispatchCommand(command bus.Command) error {
	return cqrs.commandResolver(command).Perform(command)
}

func (cqrs *busCqrs) AsyncDispatchCommand(command bus.Command) {
	cqrs.event.Emit(DISPATCH_COMMAND, command)
	log.Default().Println("Dispatch async command:" + command.Name() + "at:" + command.Timestamp())
}

func (cqrs *busCqrs) DispatchEvent(event bus.Event) error {
	return cqrs.eventResolver(event).Perform(event)
}

func (cqrs *busCqrs) DispatchQuery(query bus.Query) (any, error) {
	return cqrs.queryResolver(query).Perform(query)
}
