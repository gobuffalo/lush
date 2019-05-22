package commands

import "fmt"

type flagSlice []string

func (f *flagSlice) String() string {
	return fmt.Sprintf("%s", []string(*f))
}

func (f *flagSlice) Set(s string) error {
	*f = append(*f, s)
	return nil
}
