package multiply

import "fmt"

func Cant(at, b interface{}) error {
	return fmt.Errorf("can not multiply %T and %T", at, b)
}
