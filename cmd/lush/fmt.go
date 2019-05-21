package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gobuffalo/lush"
	"github.com/google/go-cmp/cmp"
)

var fmtOptions = struct {
	Flags *flag.FlagSet
	Write bool
	Diffs bool
}{}

func init() {
	f := flag.NewFlagSet("fmt", flag.ExitOnError)
	f.BoolVar(&fmtOptions.Write, "w", false, "write result to (source) file instead of stdout")
	f.BoolVar(&fmtOptions.Diffs, "d", false, "display diffs instead of rewriting files")
	fmtOptions.Flags = f
}

const fmtUsage = `
usage: lush fmt [-w] [-d] [files]
`

func format(args []string) {
	if len(args) == 0 {
		fmt.Println(strings.TrimSpace(fmtUsage))
		os.Exit(1)
	}
	for _, a := range args {
		script, err := lush.ParseFile(a)
		if err != nil {
			log.Fatal(err)
		}

		if fmtOptions.Diffs {
			b, err := ioutil.ReadFile(a)
			if err != nil {
				log.Fatal(err)
			}
			diff := cmp.Diff(string(b), script.String())
			fmt.Print(diff)
			continue
		}

		if !fmtOptions.Write {
			fmt.Println(script.String())
			continue
		}

		f, err := os.Create(a)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.WriteString(script.String())
		if err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
