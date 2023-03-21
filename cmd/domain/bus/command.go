package bus

type Command interface {
	Type() any
	Name() string
	Data() any
	Timestamp() string
}
