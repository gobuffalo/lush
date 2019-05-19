package ast

import (
	"encoding/json"
	"fmt"
	"go/build"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func DebugWriter(i interface{}, w io.Writer) error {
	// env := os.Getenv("debug")
	// if !testing.Verbose() || env == "debug" {
	// 	return nil
	// }
	if i == nil {
		panic("spoonman")
	}
	pwd, _ := os.Getwd()
	_, cl, l, _ := runtime.Caller(2)
	cl = strings.TrimPrefix(cl, pwd)

	for _, s := range build.Default.SrcDirs() {
		cl = strings.TrimPrefix(cl, s)
	}

	cl = strings.TrimPrefix(cl, string(filepath.Separator))

	if mi, ok := i.(withMeta); ok {
		m := Meta{
			Filename: cl,
			Line:     l,
			Original: fmt.Sprintf("%T", i),
		}
		i = mi.withMeta(m)
	}

	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
		return err
	}
	fmt.Printf("### ast/debug.go:45 string(b) (%T) -> %q %v\n", string(b), string(b), string(b))
	w.Write(b)
	w.Write([]byte("\n"))
	return nil
}

func Debug(i interface{}) error {
	// w = os.Stdout
	return DebugWriter(i, os.Stdout)
}
