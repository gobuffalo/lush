package faces

type Bool interface {
	Bool() bool
}

type Booler bool

func (b Booler) Bool() bool {
	return bool(b)
}
