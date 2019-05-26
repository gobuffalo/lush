package faces

// Bool ...
type Bool interface {
	Bool() bool
}

// Int ...
type Int interface {
	Int() int
}

// Float ...
type Float interface {
	Float() float64
}

// Slice ...
type Slice interface {
	Slice() []interface{}
}

// Map ...
type Map interface {
	Map() map[interface{}]interface{}
}
