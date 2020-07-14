package ast

import (
	"encoding/json"
	"fmt"
)

type ASTMarshaler interface {
	MarshalJSON() ([]byte, error)
}

// %#v - GoStringer
// %+v - LushStringer
// %v - Value
// %s - Stringer

func format(i fmt.Stringer, st fmt.State, verb rune) {
	switch verb {
	case 'v':
		if st.Flag('#') {
			if gs, ok := i.(fmt.GoStringer); ok {
				fmt.Fprint(st, gs.GoString())
				return
			}
		}
		if st.Flag('+') {
			if gs, ok := i.(LushStringer); ok {
				fmt.Fprint(st, gs.LushString())
				return
			}
		}
	case 'q':
		fmt.Fprintf(st, "%q", i.String())
		return
	}
	fmt.Fprint(st, i.String())
}

func genericJSON(i interface{}) map[string]interface{} {
	t := fmt.Sprintf("%T", i)
	return map[string]interface{}{
		t: i,
	}
}

func toJSON(t Node, i interface{}) ([]byte, error) {
	m := map[string]interface{}{
		fmt.Sprintf("%T", t): i,
	}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}
	return b, nil
}
