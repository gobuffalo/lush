package ast

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func NewString(b []byte) (String, error) {
	t := string(b)
	st := String{
		Original:    t,
		QuoteFormat: "%q",
	}
	if strings.HasPrefix(t, "`") {
		t = strings.TrimPrefix(t, "`")
		t = strings.TrimSuffix(t, "`")
		st.Original = t
		st.QuoteFormat = "`%s`"
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
	Original    string
	QuoteFormat string
	Meta        Meta
}

func (s String) String() string {
	return fmt.Sprintf(s.QuoteFormat, s.Original)
}

func (s String) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		if st.Flag('+') {
			b, err := toJSON(s)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				return
			}
			io.WriteString(st, b)
			return
		}
		fallthrough
	case 's':
		io.WriteString(st, s.String())
	case 'q':
		fmt.Fprintf(st, "`%q`", s.Original)
	}
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
