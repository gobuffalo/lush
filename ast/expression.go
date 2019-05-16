package ast

type Expression interface {
	Statement
	Bool(c *Context) (bool, error)
}
