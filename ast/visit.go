package ast

type Visitable interface {
	Visit(*Context) (interface{}, error)
}

type VisitableNode interface {
	Visitable
	Node
}
