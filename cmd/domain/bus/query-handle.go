package bus

type QueryHandle interface {
	Perform(Query) (any, error)
}
