package eventemitter

import "sync"

type EventEmitter struct {
	mutex  sync.Mutex
	events map[string][]chan any
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		events: make(map[string][]chan any),
	}
}

func (e *EventEmitter) On(eventName string, handler func(args ...any)) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if _, ok := e.events[eventName]; !ok {
		e.events[eventName] = make([]chan any, 0)
	}

	channel := make(chan any)
	e.events[eventName] = append(e.events[eventName], channel)

	go func() {
		for {
			data := <-channel
			handler(data)
		}
	}()
}

func (e *EventEmitter) Emit(eventName string, data any) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if channels, ok := e.events[eventName]; ok {
		for _, channel := range channels {
			channel <- data
		}
	}
}
