package faces

// Add defines an interface to support the
// "adding" of one type to another.
type Add interface {
	Add(b interface{}) (interface{}, error)
}

// Sub defines an interface to support the
// "subtracting" of one type from an another.
type Sub interface {
	Sub(b interface{}) (interface{}, error)
}

// Divide defines an interface to support the
// "dividing" of one type from an another.
type Divide interface {
	Divide(b interface{}) (interface{}, error)
}

// Equal defines an interface to support the
// checking "equality" of one type to an another.
type Equal interface {
	Equal(b interface{}) (bool, error)
}

// GreaterThan ...
type GreaterThan interface {
	GreaterThan(b interface{}) (bool, error)
}

// GreaterThanEqualTo ...
type GreaterThanEqualTo interface {
	GreaterThanEqualTo(interface{}) (bool, error)
}

// LessThan ...
type LessThan interface {
	LessThan(b interface{}) (bool, error)
}

// LessThanEqualTo ...
type LessThanEqualTo interface {
	LessThanEqualTo(interface{}) (bool, error)
}

// Match defines an interface to support the
// "matches" of one type from an another.
type Match interface {
	Match(string) (bool, error)
}

// Modulus defines an interface to support the
// take the "modulus" of one type from an another.
type Modulus interface {
	Modulus(b interface{}) (int, error)
}

// Multiply defines an interface to support the
// "multiplication" of one type with an another.
type Multiply interface {
	Multiply(b interface{}) (interface{}, error)
}

// NotEqual defines an interface to support the
// checking "equality" of one type to an another.
type NotEqual interface {
	NotEqual(b interface{}) (bool, error)
}
