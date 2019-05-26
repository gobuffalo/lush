package divide

import "fmt"

func Cant(at, b interface{}) error {
	return fmt.Errorf("can not divide %T and %T", at, b)
}
