package bus

type Command interface {
	Type() interface{}
	Name() string
	Data() interface{}
	Timestamp() string
}
