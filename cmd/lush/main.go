package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
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
	default:
		run(args)
	}
}

func run(args []string) {
	for _, a := range args {
		script, err := lush.ParseFile(a)
		if err != nil {
			log.Fatal(err)
		}
		c := ast.NewContext(context.Background(), os.Stdout)

		res, err := script.Exec(c)
		if err != nil {
			log.Fatal(err)
		}

		if res == nil {
			return
		}

		if ri, ok := res.Value.([]interface{}); ok {
			for _, i := range ri {
				fmt.Println(i)
			}
			continue
		}
		fmt.Println(res)
	}
}
