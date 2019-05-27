package commands

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
)

type Runner struct {
	*flag.FlagSet
	Args []string
}

func NewRunner() *Runner {
	r := &Runner{}
	f := flag.NewFlagSet("run", flag.ExitOnError)

	r.FlagSet = f

	return r
}

func (r Runner) Exec() error {
	for _, a := range r.FlagSet.Args() {
		script, err := lush.ParseFile(a)
		if err != nil {
			log.Fatal(err)
		}

		c := ast.NewContext(context.Background(), os.Stdout)
		builtins.Available.Range(func(k, v interface{}) bool {
			c.Imports.Store(k, v)
			return true
		})
		res, err := script.Exec(c)
		if err != nil {
			return err
		}

		if res == nil {
			return nil
		}

		if ri, ok := res.Value.([]interface{}); ok {
			for _, i := range ri {
				fmt.Println(i)
			}
			continue
		}
		fmt.Println(res)
	}
	return nil
}
