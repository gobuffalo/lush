package ast_test

import (
	"context"
	"html/template"
	"strings"
	"testing"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/stretchr/testify/require"
)

type legacyContext interface {
	context.Context
	New() legacyContext
	Has(key string) bool
	Set(key string, value interface{})
}

type legacyHelperContext interface {
	legacyContext
	Block() (string, error)
	BlockWith(legacyContext) (string, error)
	HasBlock() bool
	Render(s string) (string, error)
}

func upcase(s string, help hctx.HelperContext) (template.HTML, error) {
	var lines []string
	if help.HasBlock() {
		x, err := help.Block()
		if err != nil {
			return "", err
		}
		lines = append(lines, strings.ToUpper(x))
	}
	lines = append(lines, strings.ToUpper(s))
	return template.HTML(strings.Join(lines, "|")), nil
}

func Test_legacyHelperContext(t *testing.T) {
	r := require.New(t)

	in := `return upcase("john"){
return "Paul"
}`

	c := NewContext()
	c.Set("upcase", upcase)

	res, err := exec(in, c)
	r.NoError(err)
	r.Equal("", res.Value)
}
