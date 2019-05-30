package commands

import (
	"fmt"
	"os"
	"strings"
)

const usage = `
Lush is a tool for managing Lush source code.

Usage:

	lush <command> [arguments]

The commands are:

	run		Executes .lush files
	fmt		lushfmt (reformat) lush sources
	print	prints the .lush in various formats
	ast		print the AST for a .lush file
`

type runner interface {
	Exec([]string) error
}

func Route(args []string) error {
	u := strings.TrimSpace(usage)
	if len(args) < 1 {
		args = append(args, "-h")
	}
	var r runner
	a := args[0]
	switch a {
	case "run", "r":
		r = NewRunner(os.Stdout)
	case "fmt":
		r = NewFmter(os.Stdout)
	case "ast", "print", "a", "p":
		r = NewPrinter(os.Stdout, a)
	case "-h":
		fmt.Println(u)
	default:
		return fmt.Errorf(u)
	}

	return r.Exec(args[1:])
}
