package ast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type irritable struct {
	ind   int
	stuff []interface{}
}

func (i *irritable) Next() interface{} {
	defer func() { i.ind++ }()
	if i.ind >= len(i.stuff) {
		return nil
	}
	return i.stuff[i.ind]
}

func Test_Iterator(t *testing.T) {
	table := []struct {
		in  string
		exp []interface{}
		err bool
	}{
		{` for x, y := range arr {
			capture(y)
		}`, []interface{}{1, 2, 3}, false},
		{` for x, y := range notIterator {
			capture(y)
		}`, []interface{}{}, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()

			var ret []interface{}
			c.Set("capture", func(i interface{}) {
				ret = append(ret, i)
			})
			c.Set("arr", &irritable{stuff: tt.exp})
			c.Set("notIterator", true)

			_, err := exec(tt.in, c)

			if tt.err {
				r.Error(err)
				return
			}

			r.NoError(err)

			r.Equal(tt.exp, ret)
		})
	}
}
