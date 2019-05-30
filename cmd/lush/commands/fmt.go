package commands

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gobuffalo/lush"
	"github.com/google/go-cmp/cmp"
)

type Fmter struct {
	io.Writer
	Flags *flag.FlagSet
	Write bool
	Diffs bool
}

func NewFmter(w io.Writer) *Fmter {
	ft := &Fmter{}
	f := flag.NewFlagSet("fmt", flag.ExitOnError)
	f.BoolVar(&ft.Write, "w", false, "write result to (source) file instead of stdout")
	f.BoolVar(&ft.Diffs, "d", false, "display diffs instead of rewriting files")
	ft.Flags = f

	return ft
}

func (f *Fmter) Exec(args []string) error {
	if err := f.Flags.Parse(args); err != nil {
		return err
	}

	args = f.Flags.Args()

	if len(args) == 0 {
		return fmt.Errorf(strings.TrimSpace(fmtUsage))
	}

	for _, a := range args {
		script, err := lush.ParseFile(a)
		if err != nil {
			return err
		}

		if f.Diffs {
			b, err := ioutil.ReadFile(a)
			if err != nil {
				return err
			}
			diff := cmp.Diff(string(b), script.String())
			fmt.Fprint(f.Writer, diff)
			continue
		}

		if !f.Write {
			fmt.Println(script.String())
			continue
		}

		f, err := os.Create(a)
		if err != nil {
			return err
		}
		_, err = f.WriteString(script.String())
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}
	return nil
}

const fmtUsage = `
usage: lush fmt [-w] [-d] [files]
`
