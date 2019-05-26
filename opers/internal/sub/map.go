package sub

import "fmt"

func Map(a map[string]interface{}, b interface{}) (map[string]interface{}, error) {
	bt := fmt.Sprintf("%s", b)
	if _, ok := a[bt]; !ok {
		return a, Cant(a, b)
	}
	delete(a, bt)
	return a, nil
}
