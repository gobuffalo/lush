package faces

type GreaterThanEqualTo interface {
	GreaterThanEqualTo(interface{}) (bool, error)
}
