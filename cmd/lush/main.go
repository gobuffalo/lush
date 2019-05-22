package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gobuffalo/lush/cmd/lush/commands"
)

const usage = `
Lush is a tool for managing Lush source code.

Usage:

	lush <command> [arguments]

The commands are:

	run		Executes .lush files
	fmt		lushfmt (reformat) lush sources
	ast		print the AST for a .lush file
`

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		args = append(args, "-h")
	}
	switch args[0] {
	case "run":
		args = args[1:]
	case "fmt":
		if err := fmtOptions.Flags.Parse(args[1:]); err != nil {
			log.Fatal(err)
		}
		format(fmtOptions.Flags.Args())
		return
	case "ast":
		printAST(args[1:])
		return
	case "-h":
		fmt.Println(strings.TrimSpace(usage))
		os.Exit(1)
	}
	run(args)
}

func run(args []string) {
	r := commands.NewRunner()
	r.FlagSet.Parse(args)
	err := r.Exec()
	if err != nil {
		log.Fatal(err)
	}
}
