package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast/internal/quick"
	"github.com/stretchr/testify/require"
)

func Test_Else_Format(t *testing.T) {
	stringv, err := jsonFixture("Else")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{`%s`, " else {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%v`, " else {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%#v`, " else {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%+v`, stringv},
		{`%q`, "\" else {\\n\\tfoo = 42\\n\\n\\tfoo := 42\\n}\""},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.ELSE)

			r.Equal(tt.out, ft)
		})
	}
}
