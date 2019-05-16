package ast

type Nil struct{}

func (i Nil) IsZero() bool {
	return true
}

func (i Nil) String() string {
	return "nil"
}

func (i Nil) Interface() interface{} {
	return nil
}

func (i Nil) Exec(c *Context) (interface{}, error) {
	return nil, nil
}

func (i Nil) Bool(c *Context) (bool, error) {
	return false, nil
}
