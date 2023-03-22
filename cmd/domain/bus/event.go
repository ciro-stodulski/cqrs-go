package bus

type Event interface {
	Type() any
	Name() string
	Data() any
	Timestamp() string
}
