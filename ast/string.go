package ast

import (
	"fmt"
	"strconv"
	"strings"
)

func NewString(b []byte) (String, error) {
	t := string(b)
	st := String{
		Original: t,
		format:   "%q",
	}
	if strings.HasPrefix(t, "`") {
		t = strings.TrimPrefix(t, "`")
		t = strings.TrimSuffix(t, "`")
		st.Original = t
		st.format = "`%s`"
		return st, nil
	}
	s, err := strconv.Unquote(t)
	if err != nil {
		return st, nil
	}
	st.Original = s

	return st, nil

}

type String struct {
	Original string
	Meta     Meta
	format   string
}

func (b *String) SetMeta(m Meta) {
	b.Meta = m
}

func (s String) String() string {
	return fmt.Sprintf(s.format, s.Original)
}

func (s String) Interface() interface{} {
	return s.Original
}

func (s String) MapKey() string {
	return s.Original
}

func (s String) Exec(c *Context) (interface{}, error) {
	return s.Original, nil
}

func (s String) Bool(c *Context) (bool, error) {
	return len(s.Original) > 0, nil
}
