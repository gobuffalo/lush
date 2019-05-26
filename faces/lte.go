package faces

// LessThanEqualTo ...
type LessThanEqualTo interface {
	LessThanEqualTo(interface{}) (bool, error)
}
