package bus

type QueryHandle interface {
	Perform(query any) (any, error)
}
