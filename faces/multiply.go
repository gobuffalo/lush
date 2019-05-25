package faces

// Multiply defines an interface to support the
// "multiplication" of one type with an another.
type Multiply interface {
	Multiply(b interface{}) (interface{}, error)
}
