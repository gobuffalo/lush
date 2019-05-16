package ast

type Continue struct {
	Meta Meta
}

func (b *Continue) SetMeta(m Meta) {
	b.Meta = m
}

func (Continue) String() string {
	return "continue"
}
