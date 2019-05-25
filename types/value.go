package types

import "fmt"

// Value ...
func Value(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
