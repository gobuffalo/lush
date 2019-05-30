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

type runner interface {
	Exec([]string) error
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		args = append(args, "-h")
	}
	var r runner
	switch args[0] {
	case "run":
		r = commands.NewRunner()
	case "fmt":
		r = commands.NewFmter()
	case "ast":
		r = commands.Printer{
			Kind: "ast",
		}
	case "print":
		r = commands.Printer{
			Kind: "print",
		}
	case "-h":
		fmt.Println(strings.TrimSpace(usage))
	default:
		fmt.Println(strings.TrimSpace(usage))
		os.Exit(1)
	}

	if err := r.Exec(args[1:]); err != nil {
		log.Fatal(err)
	}
}
