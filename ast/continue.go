package ast

type Continue struct {
	Meta Meta
}

func (Continue) String() string {
	return "continue"
}
