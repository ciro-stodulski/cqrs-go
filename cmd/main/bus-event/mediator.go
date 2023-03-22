package busevent

import "cqrs-go/cmd/domain/bus"

func (cqrs *eventBus) eventResolver(command bus.Event) bus.EventHandle {
	return cqrs.containerEvent.ResolveEvent(command.Name(), command.Type())
}
