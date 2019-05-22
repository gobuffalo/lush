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
	Imports flagSlice
	Args    []string
}

func NewRunner() *Runner {
	r := &Runner{}
	f := flag.NewFlagSet("run", flag.ExitOnError)
	f.Var(&r.Imports, "import", "allows for importing of the specified package")

	r.FlagSet = f

	return r
}

func (r Runner) Exec() error {
	c := ast.NewContext(context.Background(), os.Stdout)
	for _, i := range r.Imports {
		imp, ok := builtins.Available.Load(i)
		if !ok {
			return fmt.Errorf("could not find import for %s", i)
		}
		c.Imports.Store(i, imp)
	}
	for _, a := range r.FlagSet.Args() {
		script, err := lush.ParseFile(a)
		if err != nil {
			log.Fatal(err)
		}

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
