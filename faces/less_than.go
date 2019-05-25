package faces

type LessThan interface {
	LessThan(b interface{}) (bool, error)
}
