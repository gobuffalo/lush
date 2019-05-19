package ast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func NewCall(root Statement, accessor Statement, args Statements, next Statements, block *Block) (Call, error) {
	if accessor == nil {
		fmt.Println("no with")
	}
	c := Call{
		Root:      root,
		Accessor:  accessor,
		Arguments: args,
		Next:      next,
		Block:     block,
	}
	return c, nil
}

type Call struct {
	Root       Statement
	Accessor   Statement
	Arguments  Statements
	Next       Statements
	Block      *Block
	Concurrent bool
	Meta       Meta
}

func (f Call) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"ast.Call": map[string]interface{}{
			// "Root": f.Root,
			"Accessor":   f.Accessor,
			"Arguments":  f.Arguments,
			"Next":       f.Next,
			"Block":      f.Block,
			"Concurrent": f.Concurrent,
			"Meta":       f.Meta,
		},
	}

	return json.MarshalIndent(m, "", "  ")
}

func (f Call) withMeta(m Meta) Statement {
	f.Meta = m
	return f
}

func (f Call) String() string {
	bb := &bytes.Buffer{}
	if f.Root != nil {
		bb.WriteString(f.Root.String())
	}

	if f.Accessor != nil {
		if f.Root != nil {
			bb.WriteString(".")
		}
		bb.WriteString(f.Accessor.String())
	}

	if f.Arguments != nil {
		bb.WriteString("(")
		var args []string
		for _, a := range f.Arguments {
			st := a.(fmt.Stringer)
			args = append(args, strings.TrimSpace(st.String()))
		}
		bb.WriteString(strings.Join(args, ", "))
		bb.WriteString(")")
	}

	if len(f.Next) > 0 {
		bb.WriteString(".")
	}

	var next []string

	for _, n := range f.Next {
		next = append(next, n.String())
	}

	bb.WriteString(strings.Join(next, "."))

	if f.Block != nil {
		bb.WriteString(f.Block.String())
	}

	return bb.String()
}

func (f Call) Exec(c *Context) (interface{}, error) {
	Debug(f)
	var err error
	var start interface{}

	if f.Root != nil {
		start, err = exec(c, f.Root)
		if err != nil {
			return nil, err
		}
		if fn, ok := f.Root.(Func); ok {
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
	}

	rv := reflect.ValueOf(start)
	switch t := f.Accessor.(type) {
	case Call:
		t.Root = Holder{Value: start}
		m := rv.MethodByName(t.Accessor.String())
		start, err = t.callFunc(c, m)
		if err != nil {
			return nil, err
		}
	case Func:
		start, err = t.mExec(c, f.Arguments...)
		if err != nil {
			return nil, err
		}
	case Ident:
		if !rv.IsValid() || rv.IsNil() {
			start, err = t.Exec(c)
			if err != nil {
				return nil, err
			}
			break
		}
		m := rv.MethodByName(t.Name)
		start, err = f.callFunc(c, m)
		if err != nil {
			return nil, err
		}
	}
	rv = reflect.ValueOf(start)

	var args []interface{}
	for _, a := range f.Arguments {
		if cl, ok := a.(Call); ok {
			cl.Root = Holder{Value: start}
			a = cl
		}
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
	switch rv.Kind() {
	case reflect.Func:
		start, err = f.callFunc(c, rv)
		if err != nil {
			return nil, err
		}
	case reflect.Struct:
	default:
	}

	for _, n := range f.Next {
		if vv, ok := start.(Holder); ok {
			start = vv.Value
		}
		switch t := n.(type) {
		case Call:
			t.Accessor = t.Root
			t.Root = Holder{Value: start}
			start, err = t.Exec(c)
			if err != nil {
				return nil, err
			}
		}
		if vv, ok := start.(Holder); ok {
			start = vv.Value
		}
	}
	return start, nil
}

func (f Call) callFunc(c *Context, m reflect.Value) (interface{}, error) {
	if !m.IsValid() {
		// return nil, f.Meta.Wrap(errors.New("invalid method call"))
		return nil, nil
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
					av := reflect.ValueOf(a)
					args = append(args, av)
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
	if len(args) != mt.NumIn() {
		mi := m.Interface()
		return nil, f.Meta.Errorf("%T expects %d arguments, got %d", mi, mt.NumIn(), len(args))
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
