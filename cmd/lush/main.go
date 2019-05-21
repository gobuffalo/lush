package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		switch args[0] {
		case "run":
			args = args[1:]
		case "fmt":
			if err := fmtCmd.Parse(args[1:]); err != nil {
				log.Fatal(err)
			}
			format(fmtCmd.Args())
			return
		case "ast":
			printAST(args[1:])
			return
		}
	}
	run(args)
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
