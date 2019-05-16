package ast

type Break struct {
	Meta Meta
}

func (Break) String() string {
	return "break"
}
