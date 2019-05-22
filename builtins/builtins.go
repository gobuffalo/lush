package builtins

import (
	"os"
	"sync"
)

var Available = func() sync.Map {
	m := sync.Map{}
	m.Store("fmt", NewFmt(os.Stdout))
	m.Store("strings", Strings{})
	m.Store("time", Time{})
	m.Store("os", OS{})
	return m
}()
