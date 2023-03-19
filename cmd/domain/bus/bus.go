package bus

type Bus interface {
	AsyncDispatchCommand(command Command)
	SyncDispatchCommand(command Command) error
	DispatchEvent(event *Event) error
	DispatchQuery(query *Query) (interface{}, error)
}
