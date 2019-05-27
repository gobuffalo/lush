package ast

func exec(c *Context, i interface{}) (interface{}, error) {
	ex, ok := i.(Visitable)
	if !ok {
		return i, nil
	}
	return ex.Visit(c)
}
