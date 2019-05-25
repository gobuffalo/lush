package types

import "fmt"

func String(i interface{}) string {
	switch s := i.(type) {
	case string:
		return s
	case fmt.Stringer:
		return s.String()
	}
	return fmt.Sprintf("%s", i)
}
