package ast

import (
	"fmt"

	"github.com/gobuffalo/lush/builtins"
)

type Import struct {
	Name string
	Meta Meta
}

func (i Import) String() string {
	return fmt.Sprintf(`import "%s"`, i.Name)
}

func (a Import) Visit(v Visitor) error {
	return v(a.Meta)
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

func (i Import) GoString() string {
	// c.Imports.LoadOrStore("fmt", builtins.Fmt{})

	x, ok := builtins.Available.Load(i.Name)
	if ok {
		s := `
%si, _ := c.Imports.LoadOrStore(%q, %#v)
%s, ok := %si.(%T)
if !ok {
	return nil, fmt.Errorf("expected %T got %%T", %si)
}
_ = %s
		`
		return fmt.Sprintf(s, i.Name, i.Name, x, i.Name, i.Name, x, x, i.Name, i.Name)
	}

	s := `
	if _, ok := c.Imports.Load(%q); ok {
		return nil, fmt.Errorf("could not find import for %q")
	}
	`

	return fmt.Sprintf(s, i.Name, i.Name)
}
