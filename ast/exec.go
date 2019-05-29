package ast

type Execable interface {
	Exec(*Context) (interface{}, error)
}

type ExecableNode interface {
	Execable
	Node
}

func exec(c *Context, i interface{}) (interface{}, error) {
	ex, ok := i.(Execable)
	if !ok {
		return i, nil
	}
	return ex.Exec(c)
}
