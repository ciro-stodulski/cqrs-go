package bus

type CommandHandler interface {
	Perform(Command) error
}
