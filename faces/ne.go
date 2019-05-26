package faces

// Equal defines an interface to support the
// checking "equality" of one type to an another.
type NotEqual interface {
	NotEqual(b interface{}) (bool, error)
}
