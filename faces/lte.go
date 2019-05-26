package faces

type LessThanEqualTo interface {
	LessThanEqualTo(interface{}) (bool, error)
}
