package bus

type EventHandle interface {
	Perform(Event) error
}
