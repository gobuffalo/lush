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
	Flags *flag.FlagSet
	Args  []string
}

func NewRunner(args []string) *Runner {
	r := &Runner{
		Args: args,
	}
	f := flag.NewFlagSet("run", flag.ExitOnError)

	r.Flags = f

	return r
}

func (r Runner) Exec() error {
	if err := r.Flags.Parse(r.Args); err != nil {
		return err
	}

	r.Args = r.Flags.Args()

	for _, a := range r.Args {
		script, err := lush.ParseFile(a)
		if err != nil {
			log.Fatal(err)
		}

		c := ast.NewRuntime(context.Background(), os.Stdout)
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
