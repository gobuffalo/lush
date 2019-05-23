package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/ast/internal/quick"
)

func main() {
	write(ast.Nil{})
	write(ast.Continue{})
	write(ast.Break{})
	write(ast.True)
	write(quick.STRING)
	write(quick.ARRAY)
	write(quick.IDENT)
	write(quick.INT)
	write(quick.FLOAT)
	write(quick.ASSIGN)
	write(quick.VAR)
	write(quick.ACCESS)
	write(quick.BLOCK)
	write(quick.COMMENT)
	write(quick.MAP)
	write(quick.CALL)
	write(quick.IF)
	write(quick.ELSE)
	write(quick.ELSEIF)
	write(quick.FOR)
	write(quick.RANGE)
	write(quick.FUNC)
	write(quick.LET)
	write(quick.OPEXPR)
	write(quick.RETURN)
	write(quick.IMPORT)
	compileGoTests()
}

func write(s interface{}) {

	name := fmt.Sprintf("%T", s)

	name = strings.TrimPrefix(name, "*")
	name = strings.TrimPrefix(name, "ast.")
	root := filepath.Join("ast", "testdata")
	os.MkdirAll(root, 0755)
	f, err := os.Create(filepath.Join(root, fmt.Sprintf("%s.json", name)))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}

func compileGoTests() {
	root := "goexamples"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".lush") {
			return nil
		}

		s, err := lush.ParseFile(path)
		if err != nil {
			return err
		}

		name := filepath.Base(path)
		name = strings.TrimSuffix(name, ".lush")

		d := struct {
			Original string
			Name     string
			GoCode   string
			File     string
		}{
			Original: s.String(),
			Name:     name,
			GoCode:   strings.TrimSpace(s.GoString()),
			File:     filepath.Base(path),
		}

		m := map[string]string{
			filepath.Join(strings.ToLower(name) + "_test.go"): goTest,
			filepath.Join(strings.ToLower(name) + ".go"):      goFile,
		}

		for k, v := range m {
			fp := filepath.Join(filepath.Dir(path), k)
			f, err := os.Create(fp)
			if err != nil {
				log.Fatal(err)
			}

			t, err := template.New(path).Parse(v)
			if err != nil {
				log.Fatal(err)
			}

			if err := t.Execute(f, d); err != nil {
				log.Fatal(err)
			}

			if err := f.Close(); err != nil {
				return err
			}

			c := exec.Command("goimports", "-w", fp)
			c.Stdin = os.Stdin
			c.Stderr = os.Stderr
			c.Stdout = os.Stdout

			if err := c.Run(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

const goTest = `
package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush"
)

func Test_{{.Name}}Exec(t *testing.T) {
	r := require.New(t)

	c := ast.NewContext(context.Background(), nil)

	s, err := lush.ParseFile("{{.File}}")
	r.NoError(err)
	r.True(Equal(c, s.Exec, {{.Name}}Exec))
}
`

const goFile = `
package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush"
)

/*
{{.Original}}
*/
func {{.Name}}Exec(c *ast.Context) (*ast.Returned, error) {
	{{.GoCode}}
}
`
