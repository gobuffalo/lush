package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/lush"
)

var fmtCmd = func() *flag.FlagSet {
	f := flag.NewFlagSet("fmt", flag.ExitOnError)
	f.Bool("w", false, "write result to (source) file instead of stdout")
	return f
}()

func format(args []string) {
	wr := fmtCmd.Lookup("w").Value.String()
	for _, a := range args {
		script, err := lush.ParseFile(a)
		if err != nil {
			log.Fatal(err)
		}
		if wr == "false" {
			fmt.Fprint(os.Stdout, script.String())
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
