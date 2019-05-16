package ast

type ElseIf struct {
	If
}

func (e ElseIf) String() string {
	s := e.If.String()
	return " else " + s
}

func NewElseIf(fi If) (ElseIf, error) {
	return ElseIf{
		If: fi,
	}, nil
}
