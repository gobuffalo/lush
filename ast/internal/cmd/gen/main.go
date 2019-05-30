package main

import (
	"bytes"
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
	"github.com/gobuffalo/lush/print/goprint"
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
			GoCode   string
			Name     string
			File     string
		}{
			Original: strings.TrimSpace(s.String()),
			Name:     name,
			File:     filepath.Base(path),
		}

		m := map[string]string{
			filepath.Join(strings.ToLower(name) + "_test.go"): goTest,
			filepath.Join(strings.ToLower(name) + ".go"):      goFile,
		}

		for k, v := range m {
			_, _ = k, v
			s, err := lush.Parse(path, []byte(d.Original))
			if err != nil {
				return err
			}
			bb := &bytes.Buffer{}

			cp := goprint.Printer{
				Context: goprint.Default.Context,
				Writer:  bb,
			}
			if err := cp.Print(s); err != nil {
				return err
			}

			d.GoCode = bb.String()

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
// Code generated by github.com/gobuffalo/lush. DO NOT EDIT.
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

var {{.Name}}BResult *ast.Returned

func Benchmark_{{.Name}}Exec_Go(t *testing.B) {
	var r *ast.Returned

	for i := 0; i < t.N; i++ {
		c := ast.NewContext(context.Background(), nil)
		c.Imports.Store("fmt", builtins.NewFmt(ioutil.Discard))

		r, _ = {{.Name}}Exec(c)
	}
	{{.Name}}BResult = r
}

func Benchmark_{{.Name}}Exec_Lush(t *testing.B) {
	var r *ast.Returned

	s, err := lush.ParseFile("{{.File}}")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < t.N; i++ {
		c := ast.NewContext(context.Background(), nil)
		c.Imports.Store("fmt", builtins.NewFmt(ioutil.Discard))

		r, _ = s.Exec(c)
	}
	{{.Name}}BResult = r
}

`

const goFile = `
// Code generated by github.com/gobuffalo/lush. DO NOT EDIT.
package goexamples

import (
	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/print/gop"
)

/*
{{.Original}}
*/
func {{.Name}}Exec(current *ast.Context) (*ast.Returned, error) {
	{{.GoCode}}
}
`
