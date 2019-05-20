package ast

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type Map struct {
	Values map[Statement]interface{}
	Meta   Meta
}

type Keyable interface {
	MapKey() string
}

func NewMap(vals interface{}) (Map, error) {
	m := Map{Values: map[Statement]interface{}{}}

	sl, err := toII(vals)
	if err != nil {
		return m, err
	}

	if len(sl) == 0 {
		return m, nil
	}

	for _, xsl := range sl {
		sl, err = toII(xsl)
		if err != nil {
			return m, err
		}

		k := sl[0]
		v := sl[4]

		sk, ok := k.(Statement)
		if !ok {
			return m, fmt.Errorf("expected Statement got %T", k)
		}
		m.Values[sk] = v
	}

	return m, nil
}

func (m Map) Exec(c *Context) (interface{}, error) {
	mm := map[interface{}]interface{}{}
	for k, v := range m.Values {
		var key interface{}
		var value interface{}

		if vv, ok := v.(interfacer); ok {
			value = vv.Interface()
		}

		if vv, ok := k.(interfacer); ok {
			key = vv.Interface()
		}

		value, err := exec(c, v)
		if err != nil {
			return nil, err
		}

		key, err = exec(c, k)
		if err != nil {
			return nil, err
		}

		mm[key] = value
	}
	return mm, nil
}

func (m Map) String() string {
	var keys []Statement

	for k := range m.Values {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(a, b int) bool {
		return keys[a].String() < keys[b].String()
	})

	bb := &bytes.Buffer{}
	bb.WriteString("{")
	var lines []string
	for _, k := range keys {
		v := m.Values[k]
		mk := strings.TrimSpace(k.String())
		lines = append(lines, fmt.Sprintf("%s: %s", mk, v))
	}
	sort.Strings(lines)
	bb.WriteString(strings.Join(lines, ", "))
	bb.WriteString("}")
	return bb.String()
}

func (m Map) Interface() interface{} {
	mm := map[string]interface{}{}

	for k, v := range m.Values {
		ks := k.String()
		if kv, ok := k.(interfacer); ok {
			ks = fmt.Sprint(kv.Interface())
		}
		if iv, ok := v.(interfacer); ok {
			v = iv.Interface()
		}
		mm[ks] = v
	}

	return mm
}

func (m Map) MarshalJSON() ([]byte, error) {
	var vals [][]interface{}

	for k, v := range m.Values {
		vals = append(vals, []interface{}{k, v})
	}

	mm := map[string]interface{}{
		"Values": vals,
	}
	return toJSON("ast.Map", mm)
}

func (m Map) Bool(c *Context) (bool, error) {
	return len(m.Values) > 0, nil
}

func (a Map) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}
