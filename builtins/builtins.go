package builtins

import (
	"os"
	"sync"
)

var Available sync.Map

var _ = func() error {
	Available.Store("fmt", NewFmt(os.Stdout))
	Available.Store("strings", Strings{})
	Available.Store("time", Time{})
	Available.Store("os", OS{})
	Available.Store("sort", Sort{})
	return nil
}()
