package ast

type Visitable interface {
	Visit(*Context) (interface{}, error)
}

type VisitableStatement interface {
	Visitable
	Statement
}
