package ast

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type value struct {
	value interface{}
}

func (v value) String() string {
	return fmt.Sprint(v.value)
}

func NewCall(name Statement, with Statement, args Statements, next Statements, block *Block) (Call, error) {
	c := Call{
		Name:      name,
		With:      with,
		Arguments: args,
		Next:      next,
		Block:     block,
	}
	return c, nil
}

type Call struct {
	Name       Statement
	With       Statement
	Arguments  Statements
	Next       Statements
	Block      *Block
	Concurrent bool
	Meta       Meta
}

func (f Call) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString(f.Name.String())
	if f.Next != nil {
		bb.WriteString(".")
		bb.WriteString(f.Next.String())
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

func (f Call) Exec(c *Context) (interface{}, error) {
	start, err := exec(c, f.Name)
	if err != nil {
		return nil, err
	}
	if fn, ok := f.Name.(Func); ok {
		start, err = fn.mExec(c, f.Arguments...)
		if err != nil {
			return nil, err
		}
	}

	if fn, ok := start.(Func); ok {
		start, err = fn.mExec(c, f.Arguments...)
		if err != nil {
			return nil, err
		}
	}

	switch t := f.With.(type) {
	case Call:
		rv := reflect.ValueOf(start)
		m := rv.MethodByName(t.Name.String())
		start, err = t.callFunc(c, m)
		if err != nil {
			return nil, err
		}
	case Func:
		panic(t)
	case Ident:
		panic(t)
	}

	var args []interface{}
	for _, a := range f.Arguments {
		i, err := exec(c, a)
		if err != nil {
			return nil, err
		}
		if s, ok := i.(interfacer); ok {
			args = append(args, s.Interface())
			continue
		}
		args = append(args, i)
	}
	rv := reflect.ValueOf(start)
	switch rv.Kind() {
	case reflect.Func:
		start, err = f.callFunc(c, rv)
		if err != nil {
			return nil, err
		}
	case reflect.Struct:
	default:
	}
	return start, nil
}

func (f Call) callFunc(c *Context, m reflect.Value) (interface{}, error) {
	if !m.IsValid() {
		return nil, f.Meta.Wrap(errors.New("invalid method call"))
	}

	mt := m.Type()

	var err error
	var args []reflect.Value
	if mt.IsVariadic() {
		for i := 0; i < len(f.Arguments); i++ {
			v := f.Arguments[i].(interface{})
			if ex, ok := v.(Execable); ok {
				v, err = ex.Exec(c)
				if err != nil {
					return nil, err
				}
			}
			if ii, ok := v.(interfacer); ok {
				v = ii.Interface()
			}
			args = append(args, reflect.ValueOf(v))
		}
	} else {
		for i := 0; i < mt.NumIn(); i++ {
			if i < len(f.Arguments) {
				v := f.Arguments[i].(interface{})
				var err error
				v, err = exec(c, v)
				if err != nil {
					return nil, err
				}
				ar := flatten([]interface{}{v})
				for _, a := range ar {
					if m, ok := a.(map[interface{}]interface{}); ok {
						mm := map[string]interface{}{}
						for k, v := range m {
							var key = fmt.Sprint(k)
							mm[key] = v
						}
						a = mm
					}
					if ii, ok := a.(interfacer); ok {
						a = ii.Interface()
					}
					args = append(args, reflect.ValueOf(a))
				}
				continue
			}
			v := mt.In(i)
			rv := reflect.Indirect(reflect.New(v))

			if _, ok := rv.Interface().(*Context); ok {
				ctx := c.Clone()
				ctx.Block = f.Block

				args = append(args, reflect.ValueOf(ctx))
				continue
			}

			if _, ok := rv.Interface().(map[string]interface{}); ok {
				args = append(args, reflect.ValueOf(map[string]interface{}{}))
			}
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
		i := res[0].Interface()
		if ii, ok := i.(interfacer); ok {
			i = ii.Interface()
		}
		return i, nil
	}

	var ins []interface{}
	for _, v := range res {
		i := v.Interface()
		if ii, ok := i.(interfacer); ok {
			i = ii.Interface()
		}
		ins = append(ins, i)
	}
	return ins, nil
}
