package faces

// Modulus defines an interface to support the
// take the "modulus" of one type from an another.
type Modulus interface {
	Modulus(b interface{}) (int, error)
}
