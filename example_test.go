package lush

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gobuffalo/lush/ast"
)

func ExampleExec() {
	c := ast.NewContext(context.Background(), os.Stdout)

	in := `return "hi"`

	res, err := ExecReader(c, "x.lush", strings.NewReader(in))
	if err != nil {
		log.Fatal(err)
	}

	if res.Value != "hi" {
		log.Fatalf("expected hi got %s", res.Value)
	}
}

func ExampleExec_assignment() {
	c := ast.NewContext(context.Background(), os.Stdout)

	in := `x := 0
	func() {
		if true {
			x = 42
		}
	}()

	fmt.Println(x)`

	_, err := ExecReader(c, "x.lush", strings.NewReader(in))
	if err != nil {
		log.Fatal(err)
	}

	// output:
	// 42
}

func ExampleExec_ifStatements() {
	c := ast.NewContext(context.Background(), os.Stdout)

	in := `
if false {
  fmt.Println("in if")
} else if (1 == 2) {
  fmt.Println("in else")
} else if true {
  fmt.Println("2 == 2")
} else {
  fmt.Println("in other else")
}
`

	_, err := ExecReader(c, "x.lush", strings.NewReader(in))
	if err != nil {
		log.Fatal(err)
	}

	// output:
	// 2 == 2
}

func ExampleExec_arrays() {
	c := ast.NewContext(context.Background(), os.Stdout)

	in := `
	a := [1, "a", true, [4, 5, nil]]
	for i, v := range a {
		fmt.Println(i, v)
	}
	`

	_, err := ExecReader(c, "x.lush", strings.NewReader(in))
	if err != nil {
		log.Fatal(err)
	}

	// output:
	// 0 1
	// 1 a
	// 2 true
	// 3 4 5
}

func ExampleExec_maps() {
	c := ast.NewContext(context.Background(), os.Stdout)

	in := `
	m := {"a": "b", "h": 1, "foo": "bar", "y": func(x) {}}
	for k, v := range m {
		fmt.Println(k, v)
	}
	`

	_, err := ExecReader(c, "x.lush", strings.NewReader(in))
	if err != nil {
		log.Fatal(err)
	}

}

func ExampleExec_infiniteForLoop() {
	c := ast.NewContext(context.Background(), os.Stdout)

	in := `for {
  if (i == nil) {
    let i = 0
  }

  i = (i + 1)

  if (i == 4) {
		fmt.Println(i)
		break
  }
}`

	_, err := ExecReader(c, "x.lush", strings.NewReader(in))
	if err != nil {
		log.Fatal(err)
	}

	// output:
	// 4

}

func ExampleExec_customHelperOptionalContext() {
	c := ast.NewContext(context.Background(), os.Stdout)
	c.Set("myFunc", func(s string, c *ast.Context) (string, error) {
		if c.Block != nil {
			res, err := c.Block.Exec(c)
			if err != nil {
				return "", err
			}
			return fmt.Sprint(res), nil
		}
		return strings.ToUpper(s), nil
	})

	in := `x := "A String"
fmt.Println(myFunc(x)) // A STRING

s := myFunc(x) {
  return "another string"
}
fmt.Println(s) // another string
`

	_, err := ExecReader(c, "x.lush", strings.NewReader(in))
	if err != nil {
		log.Fatal(err)
	}
	// output:
	// A STRING
	// another string
}
