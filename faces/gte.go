package faces

// GreaterThanEqualTo ...
type GreaterThanEqualTo interface {
	GreaterThanEqualTo(interface{}) (bool, error)
}
