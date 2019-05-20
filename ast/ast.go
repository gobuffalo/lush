package ast

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func genericJSON(i interface{}) map[string]interface{} {
	fmt.Printf("### ast/ast.go:11 i (%T) -> %q %+v\n", i, i, i)
	t := fmt.Sprintf("%T", i)
	return map[string]interface{}{
		t: i,
	}
}

func toJSON(t string, i interface{}) ([]byte, error) {
	m := map[string]interface{}{
		t: i,
	}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func printV(st fmt.State, i interface{}) {
	if !st.Flag('+') {
		if s, ok := i.(fmt.Stringer); ok {
			io.WriteString(st, s.String())
		}
		return
	}

	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	st.Write(b)
}
