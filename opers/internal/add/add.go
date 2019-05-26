package add

import "fmt"

func Cant(at, b interface{}) error {
	return fmt.Errorf("can not Add %T and %T", at, b)
}
