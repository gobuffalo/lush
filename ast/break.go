package ast

type Break struct {
	Meta Meta
}

func (b *Break) SetMeta(m Meta) {
	b.Meta = m
}

func (Break) String() string {
	return "break"
}
