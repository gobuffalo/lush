package ast

// VarRef represents the reference of a variable
type VarRef struct {
	Name string // the name of the variable being accessed
}

// Exec retrieves the named variable from the current lexical scope referred to
// by the passed in context
func (v VarRef) Exec(c *Context) (interface{}, error) {
	return c.Value(v.Name), nil
}
