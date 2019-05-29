package ast

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gobuffalo/lush/types"
)

func NewCall(n Node, y interface{}, args Nodes, b *Block) (Call, error) {
	c := Call{
		Name:      n,
		Arguments: args,
		Block:     b,
	}

	if y != nil {
		in, ok := y.(Ident)
		if !ok {
			return Call{}, fmt.Errorf("expected %T to be Ident", y)
		}
		c.FName = in
	}

	return c, nil
}

type Call struct {
	Name      Node
	FName     Ident
	Arguments Nodes
	Block     *Block
	Meta      Meta
}

func (f Call) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Name":      f.Name,
		"FName":     f.FName,
		"Arguments": f.Arguments,
		"Block":     f.Block,
		"Meta":      f.Meta,
	}

	return toJSON(f, m)
}

func (f Call) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString(f.Name.String())
	if (f.FName != Ident{}) {
		bb.WriteString(".")
		bb.WriteString(f.FName.String())
	}
	bb.WriteString("(")
	var args []string
	for _, a := range f.Arguments {
		st := a.(fmt.Stringer)
		args = append(args, strings.TrimSpace(st.String()))
	}
	bb.WriteString(strings.Join(args, ", "))
	bb.WriteString(")")
	return bb.String()
}

func (f Call) Exec(c *Runtime) (interface{}, error) {
	return f.exec(c)
}

func (f Call) exec(c *Runtime) (interface{}, error) {
	n, err := exec(c, f.Name)
	if err != nil {
		return nil, err
	}
	rv := reflect.Indirect(reflect.ValueOf(n))
	if !f.FName.IsZero() {
		m := rv.MethodByName(f.FName.String())
		return f.mExec(m, c)
	}

	return f.mExec(rv, c)
}

func (f Call) mExec(m reflect.Value, c *Runtime) (interface{}, error) {
	if !m.IsValid() {
		return nil, f.Meta.Wrap(errors.New("invalid method call"))
	}

	if fun, ok := m.Interface().(Func); ok {
		c = c.Clone()
		return fun.mExec(c, f.Arguments...)
	}

	var args []reflect.Value
	mt := m.Type()
	var err error
	if mt.IsVariadic() {
		for i := 0; i < len(f.Arguments); i++ {
			v := f.Arguments[i].(interface{})
			if args, err = app(args, mt, 0, c, v); err != nil {
				return nil, err
			}
		}
	} else {
		for i := 0; i < mt.NumIn(); i++ {
			if i < len(f.Arguments) {
				v := f.Arguments[i].(interface{})
				if args, err = app(args, mt, i, c, v); err != nil {
					return nil, err
				}
				continue
			}
			v := mt.In(i)
			rv := reflect.Indirect(reflect.New(v))

			if _, ok := rv.Interface().(*Runtime); ok {
				ctx := c.Clone()
				ctx.Block = f.Block
				args = append(args, reflect.ValueOf(ctx))
				continue
			}

			if _, ok := rv.Interface().(map[string]interface{}); ok {
				args = append(args, reflect.ValueOf(map[string]interface{}{}))
			}
			continue
		}
	}

	res := m.Call(args)
	if len(res) == 0 {
		return nil, nil
	}
	if len(res) > 0 {
		if e, ok := res[len(res)-1].Interface().(error); ok {
			return nil, e
		}
		return res[0].Interface(), nil
	}

	var ins []interface{}
	for _, v := range res {
		ins = append(ins, v.Interface())
	}
	return ins, nil
}

func app(args []reflect.Value, mt reflect.Type, i int, c *Runtime, v interface{}) ([]reflect.Value, error) {
	if m, ok := v.(Map); ok {
		args = append(args, reflect.ValueOf(m.Interface()))
		return args, nil
	}
	if ex, ok := v.(Execable); ok {
		x, err := ex.Exec(c)
		if err != nil {
			return args, err
		}
		v = x
	}
	if vi, ok := v.(interfacer); ok {
		v = vi.Interface()
	}

	app := func(v interface{}) {
		var ar reflect.Value
		expectedT := mt.In(i)
		if v != nil {
			ar = reflect.ValueOf(v)
		} else {
			ar = reflect.New(expectedT).Elem()
		}

		args = append(args, ar)
	}

	for _, x := range types.Slice(v) {
		app(x)
	}
	return args, nil
}

func (f Call) Format(st fmt.State, verb rune) {
	format(f, st, verb)
}
