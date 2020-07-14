package ast

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gobuffalo/lush/types"
)

type Map struct {
	Values map[string]interface{}
	Meta   Meta
}

func (m Map) Map() map[string]interface{} {
	mm := map[string]interface{}{}

	for k, v := range m.Values {
		if iv, ok := v.(interfacer); ok {
			v = iv.Interface()
		}
		mm[k] = v
	}

	return mm
}

func NewMap(vals interface{}) (Map, error) {
	m := Map{Values: map[string]interface{}{}}

	sl := types.Slice(vals)

	if len(sl) == 0 {
		return m, nil
	}

	for _, xsl := range sl {
		sl = types.Slice(xsl)

		k := sl[0]
		v := sl[4]
		sk := fmt.Sprintf("%s", k)
		if uq, err := strconv.Unquote(sk); err == nil {
			sk = uq
		}
		m.Values[sk] = v
	}

	return m, nil
}

func (m Map) Exec(c *Context) (interface{}, error) {
	mm := map[string]interface{}{}
	for k, v := range m.Values {
		if vv, ok := v.(interfacer); ok {
			v = vv.Interface()
		}

		value, err := exec(c, v)
		if err != nil {
			return nil, err
		}

		mm[k] = value
	}
	return mm, nil
}

func (m Map) GoString() string {
	return fmt.Sprintf("%#v", m.Values)
}

func (m Map) String() string {
	var keys []string

	for k := range m.Values {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(a, b int) bool {
		return keys[a] < keys[b]
	})

	bb := &bytes.Buffer{}
	bb.WriteString("{")
	var lines []string
	for _, k := range keys {
		v := m.Values[k]
		mk := strings.TrimSpace(k)
		lines = append(lines, fmt.Sprintf("%q: %s", mk, v))
	}
	sort.Strings(lines)
	bb.WriteString(strings.Join(lines, ", "))
	bb.WriteString("}")
	return bb.String()
}

func (m Map) Interface() interface{} {
	return m.Map()
}

func (m Map) MarshalJSON() ([]byte, error) {
	var vals [][]interface{}

	for k, v := range m.Values {
		vals = append(vals, []interface{}{k, v})
	}

	mm := map[string]interface{}{
		"Values": vals,
	}
	return toJSON(m, mm)
}

func (m Map) Bool(c *Context) (bool, error) {
	return len(m.Values) > 0, nil
}

func (a Map) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}
