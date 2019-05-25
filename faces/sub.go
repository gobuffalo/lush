package faces

// Sub defines an interface to support the
// "subtracting" of one type from an another.
type Sub interface {
	Sub(b interface{}) (interface{}, error)
}
