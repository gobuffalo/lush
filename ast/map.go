package ast

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type Map map[Statement]interface{}

type Keyable interface {
	MapKey() string
}

func NewMap(vals interface{}) (Map, error) {
	m := Map{}

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
		m[sk] = v
	}

	return m, nil
}

func (m Map) Exec(c *Context) (interface{}, error) {
	mm := map[interface{}]interface{}{}
	for k, v := range m {
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
			key = k
		}

		mm[key] = value
	}
	return mm, nil
}

func (m Map) String() string {
	var keys []Statement

	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(a, b int) bool {
		return keys[a].String() < keys[b].String()
	})

	bb := &bytes.Buffer{}
	bb.WriteString("{")
	var lines []string
	for _, k := range keys {
		v := m[k]
		mk := strings.TrimSpace(k.String())
		mv := strings.TrimSpace(fmt.Sprint(v))
		lines = append(lines, fmt.Sprintf("%s: %s", mk, mv))
	}
	sort.Strings(lines)
	bb.WriteString(strings.Join(lines, ", "))
	bb.WriteString("}")
	return bb.String()
}

func (m Map) Interface() interface{} {
	mm := map[string]interface{}{}

	for k, v := range m {
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

func (m Map) Bool(c *Context) (bool, error) {
	return len(m) > 0, nil
}

func (m Map) Equal(m2 Map) bool {
	return fmt.Sprint(m.Interface()) == fmt.Sprint(m.Interface())
}
