package ast

type Visitable interface {
	Visit(*Context) (interface{}, error)
}

type VisitableStatement interface {
	Visitable
	Statement
}

func exec(c *Context, i interface{}) (interface{}, error) {
	ex, ok := i.(Visitable)
	if !ok {
		return i, nil
	}
	return ex.Visit(c)
}
