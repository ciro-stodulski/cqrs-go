package bus

type Handler interface {
	Perform(Command) error
}
