package builtins

import "sort"

type Sort struct{}

func (Sort) GoString() string {
	return "builtins.Sort{}"
}

func (Sort) Strings(s []string) {
	sort.Strings(s)
}

func (Sort) Sort(s sort.Interface) {
	sort.Sort(s)
}
