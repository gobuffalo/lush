package ast

type Expression interface {
	Node
	Bool(c *Context) (bool, error)
}
