package add

import "fmt"

func Cant(at, b interface{}) error {
	return fmt.Errorf("can not add %T and %T", at, b)
}
