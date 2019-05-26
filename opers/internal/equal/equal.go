package equal

import "fmt"

func Cant(a interface{}, pattern string) error {
	return fmt.Errorf("can not compare %T and %s", a, pattern)
}
