package types

// Booler ...
type Booler bool

// Bool ...
func (b Booler) Bool() bool {
	return bool(b)
}
