package faces

// Bool ...
type Bool interface {
	Bool() bool
}

// Booler ...
type Booler bool

// Bool ...
func (b Booler) Bool() bool {
	return bool(b)
}
