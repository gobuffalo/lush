package builtins

import "os"

type OS struct{}

func (OS) Environ() []string {
	return os.Environ()
}

func (OS) GoString() string {
	return "builtins.OS{}"
}
