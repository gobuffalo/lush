package faces

// GreaterThan ...
type GreaterThan interface {
	GreaterThan(b interface{}) (bool, error)
}
