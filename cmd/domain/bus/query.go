package bus

type Query interface {
	Name() string
	Timestamp() string
	Data() any
	Type() any
}
