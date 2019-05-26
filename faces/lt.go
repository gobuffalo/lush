package faces

// LessThan ...
type LessThan interface {
	LessThan(b interface{}) (bool, error)
}
