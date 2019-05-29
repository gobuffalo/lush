package ast

type Execable interface {
	Exec(*Runtime) (interface{}, error)
}

type ExecableNode interface {
	Execable
	Node
}

func exec(c *Runtime, i interface{}) (interface{}, error) {
	ex, ok := i.(Execable)
	if !ok {
		return i, nil
	}
	return ex.Exec(c)
}
