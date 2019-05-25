package faces

// Divide defines an interface to support the
// "dividing" of one type from an another.
type Divide interface {
	Divide(b interface{}) (interface{}, error)
}
