package ast

type Script struct {
	Statements Statements
}

func (s Script) Exec(c *Context) (*Returned, error) {
	res, err := s.Statements.Exec(c)
	if err != nil {
		return nil, err
	}

	c.wg.Wait()

	ret, ok := res.(Returned)
	if !ok {
		return nil, nil
	}
	return &ret, nil
}

func (s Script) String() string {
	return s.Statements.String()
}
