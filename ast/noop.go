package ast

type Noop struct {
	Text string
}

func (n Noop) String() string {
	return ""
}

func NewNoop(b []byte) (Noop, error) {
	return Noop{
		Text: string(b),
	}, nil
}
