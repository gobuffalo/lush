package ast

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func toJSON(i interface{}) (string, error) {
	m := map[string]interface{}{}
	m[fmt.Sprintf("%T", i)] = i
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func printV(st fmt.State, i interface{}) {
	if !st.Flag('+') {
		if s, ok := i.(fmt.Stringer); ok {
			io.WriteString(st, s.String())
		}
		return
	}

	b, err := toJSON(i)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	io.WriteString(st, b)
}
