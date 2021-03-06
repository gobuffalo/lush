package ast

import (
	"fmt"
)

type Import struct {
	Name string
	Meta Meta
}

func (i Import) String() string {
	return fmt.Sprintf(`import "%s"`, i.Name)
}

func (i Import) Exec(c *Context) (interface{}, error) {
	imp, ok := c.Imports.Load(i.Name)
	if !ok {
		return nil, fmt.Errorf("could not find import for %s", i.Name)
	}
	c.Set(i.Name, imp)
	return nil, nil
}

func (i Import) Format(st fmt.State, verb rune) {
	format(i, st, verb)
}

func (i Import) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Name": genericJSON(i.Name),
		"Meta": i.Meta,
	}
	return toJSON(i, m)
}

func NewImport(s string) (Import, error) {
	return Import{Name: s}, nil
}
