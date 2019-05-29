package ast

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func NewFor(n ExecableNode, args interface{}, b *Block) (For, error) {
	f := For{
		Block: b,
		Name:  n,
	}

	if args == nil {
		return f, nil
	}

	ii := flatten([]interface{}{args})

	if len(ii) == 0 {
		return f, errors.New("for requires at least one var name")
	}
	if len(ii) > 2 {
		return f, errors.New("for can not take more than 2 variables")
	}

	for _, a := range ii {
		i, ok := a.(Ident)
		if !ok {
			return f, fmt.Errorf("expected Ident, got %T", a)
		}
		f.Args = append(f.Args, i)
	}

	return f, nil
}

type For struct {
	Name         ExecableNode
	Args         Idents
	Block        *Block
	Meta         Meta
	normalSingle bool
}

func (f For) Format(st fmt.State, verb rune) {
	format(f, st, verb)
}

func (f For) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Name":         f.Name,
		"Args":         f.Args,
		"Block":        f.Block,
		"Meta":         f.Meta,
		"NormalSingle": f.normalSingle,
	}
	return toJSON(f, m)
}

func (f For) String() string {
	if f.Name == nil {
		bb := &bytes.Buffer{}
		bb.WriteString("for ")
		if f.Block != nil {
			bb.WriteString(f.Block.String())
		}
		return bb.String()
	}
	bb := &bytes.Buffer{}
	bb.WriteString("for (")
	var args []string
	for _, a := range f.Args {
		args = append(args, a.String())
	}
	bb.WriteString(strings.Join(args, ", "))
	bb.WriteString(") in ")
	bb.WriteString(f.Name.String() + " ")
	if f.Block != nil {
		bb.WriteString(f.Block.String())
	}
	return bb.String()
}

func (f For) Exec(c *Runtime) (interface{}, error) {
	c = c.Clone()

	var v interface{}
	var err error
	if f.Name != nil {
		v, err = f.Name.Exec(c)
		if err != nil {
			return nil, err
		}
	}

	if v == nil {
		for {
			res, err := f.Block.Exec(c)
			if err != nil {
				return nil, err
			}
			switch rr := res.(type) {
			case Break:
				return nil, nil
			case Return, Returned:
				return rr, nil
			case Continue:
				continue
			}
		}
	}

	if it, ok := v.(Iterator); ok {
		n := it.Next()
		var i int
		for n != nil {
			rv := reflect.ValueOf(n)
			res, err := f.iterate(c, rv, i, n)
			if err != nil {
				return nil, err
			}
			switch rr := res.(type) {
			case Returned:
				return rr, nil
			case Return, Break:
				return rr, nil
			case Continue:
				continue
			}
			n = it.Next()
			i++
		}
		return nil, nil
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			res, err := f.iterate(c, rv, i, rv.Index(i))
			if err != nil {
				return nil, err
			}
			switch rr := res.(type) {
			case Returned:
				return rr, nil
			case Return, Break:
				return rr, nil
			case Continue:
				continue
			}
		}
		return nil, nil
	case reflect.Map:
		for _, k := range rv.MapKeys() {
			res, err := f.iterate(c, rv, k, rv.MapIndex(k))
			if err != nil {
				return nil, err
			}
			switch rr := res.(type) {
			case Return, Break:
				return rr, nil
			case Continue:
				continue
			}
		}
		return nil, nil
	default:

	}
	return nil, f.Meta.Errorf("can't iterate over %T", v)
}

func (f For) iterate(c *Runtime, rv reflect.Value, k interface{}, v interface{}) (interface{}, error) {
	if len(f.Args) == 1 {
		if !f.normalSingle {
			k = v
		}
		c.Set(f.Args[0].String(), k)
	}
	if len(f.Args) == 2 {
		c.Set(f.Args[0].String(), k)
		c.Set(f.Args[1].String(), v)
	}
	if f.Block == nil {
		return nil, nil
	}
	r, err := f.Block.Exec(c)
	if err != nil {
		return nil, err
	}
	return r, nil
}
