package builtins

import "os"

var OS = bos{
	Args: os.Args,
}

type bos struct {
	Args []string
}

func (bos) Environ() []string {
	return os.Environ()
}

func (bos) GoString() string {
	return "builtins.OS"
}
