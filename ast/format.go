package ast

import (
	"encoding/json"
	"fmt"
	"os"
)

type ASTMarshaler interface {
	MarshalJSON() ([]byte, error)
}

func format(i fmt.Stringer, st fmt.State, verb rune) {
	switch verb {
	case 'v':
		if st.Flag('+') {
			b, err := json.MarshalIndent(i, "", "  ")
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				return
			}
			fmt.Fprint(st, string(b))
			return
		}
		fmt.Fprint(st, i.String())
	case 'q':
		fmt.Fprintf(st, "%q", i.String())
	default:
		fmt.Fprint(st, i.String())
	}
}

func genericJSON(i interface{}) map[string]interface{} {
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
