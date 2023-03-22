package bus

type Bus interface {
	AsyncDispatchCommand(command Command)
	SyncDispatchCommand(command Command) error
	DispatchEvent(event Event) error
	DispatchQuery(query Query) (any, error)
}

type EventBus interface {
	DispatchEvent(event Event) error
}
