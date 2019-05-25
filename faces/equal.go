package faces

// Equal defines an interface to support the
// checking "equality" of one type to an another.
type Equal interface {
	Equal(b interface{}) (bool, error)
}
