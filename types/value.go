package types

import "fmt"

func Value(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
