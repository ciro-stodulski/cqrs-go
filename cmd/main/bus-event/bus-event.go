package busevent

import (
	"cqrs-go/cmd/domain/bus"
)

type eventBus struct {
	containerEvent *containerEvent
}

func NewEventBus() bus.EventBus {
	return &eventBus{containerEvent: newContainerEvent()}
}

func (cqrs *eventBus) DispatchEvent(event bus.Event) error {
	return cqrs.eventResolver(event).Perform(event)
}
