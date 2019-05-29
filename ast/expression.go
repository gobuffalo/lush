package ast

type Expression interface {
	Node
	Bool(c *Runtime) (bool, error)
}
