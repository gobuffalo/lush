package commands

import (
	"encoding/json"
	"fmt"

	"github.com/gobuffalo/lush"
)

const astUsage = `
usage: lush ast [files]
`

type AstPrinter struct{}

func (a AstPrinter) Exec(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must pass at a least one argument")
	}
	for _, a := range args {
		script, err := lush.ParseFile(a)
		if err != nil {
			return err
		}
		b, err := json.MarshalIndent(script, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	}
	return nil
}
