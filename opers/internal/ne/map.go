package ne

import (
	"fmt"
	"sort"
	"strings"
)

// Map ...
func Map(m map[interface{}]interface{}, b interface{}) (bool, error) {
	switch bt := b.(type) {
	case map[interface{}]interface{}:
		var sa []string
		var sb []string

		for k, v := range m {
			sa = append(sa, fmt.Sprintf("%v=%v", k, v))
		}
		for k, v := range bt {
			sb = append(sb, fmt.Sprintf("%v=%v", k, v))
		}

		sort.Strings(sa)
		sort.Strings(sb)

		return strings.Join(sa, ",") != strings.Join(sb, ","), nil
	}
	return false, Cant(m, fmt.Sprintf("%T", b))
}
