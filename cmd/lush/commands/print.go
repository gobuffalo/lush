package commands

import (
	"encoding/json"
	"fmt"

	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/print/goprint"
)

type Printer struct {
	Kind string
}

func (a Printer) Exec(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must pass at a least one argument")
	}
	switch a.Kind {
	case "ast":
		return a.astExec(args)
	case "print":
		if len(args) == 0 {
			return fmt.Errorf("you must pass at a least one argument and a language")
		}
		switch args[0] {
		case "go":
			return a.goExec(args[1:])
		case "ast":
			return a.astExec(args[1:])
		case "lush":
			return a.lushExec(args[1:])
		}
	}
	return fmt.Errorf("unknown printer kind %q", a.Kind)
}

func (a Printer) lushExec(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must pass at a least one argument")
	}
	for _, a := range args {
		script, err := lush.ParseFile(a)
		if err != nil {
			return err
		}
		fmt.Println(script)
	}
	return nil
}

func (a Printer) goExec(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must pass at a least one argument")
	}
	for _, a := range args {
		script, err := lush.ParseFile(a)
		if err != nil {
			return err
		}
		err = goprint.Default.Print(script.Nodes...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a Printer) astExec(args []string) error {
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
