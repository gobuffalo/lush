package faces

// Add defines an interface to support the
// "adding" of one type to another.
type Add interface {
	Add(b interface{}) (interface{}, error)
}
