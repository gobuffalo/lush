package commands

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gobuffalo/lush"
	"github.com/google/go-cmp/cmp"
)

type Fmter struct {
	Flags *flag.FlagSet
	Args  []string
	Write bool
	Diffs bool
}

func NewFmter(args []string) *Fmter {
	ft := &Fmter{
		Args: args,
	}
	f := flag.NewFlagSet("fmt", flag.ExitOnError)
	f.BoolVar(&ft.Write, "w", false, "write result to (source) file instead of stdout")
	f.BoolVar(&ft.Diffs, "d", false, "display diffs instead of rewriting files")
	ft.Flags = f

	return ft
}

func (f *Fmter) Exec() error {
	if err := f.Flags.Parse(f.Args); err != nil {
		return err
	}

	f.Args = f.Flags.Args()

	if len(f.Args) == 0 {
		fmt.Println(strings.TrimSpace(fmtUsage))
		os.Exit(1)
	}
	for _, a := range f.Args {
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
			fmt.Print(diff)
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
