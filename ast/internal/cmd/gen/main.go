package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

		fp := filepath.Join(filepath.Dir(path), strings.ToLower(name)+".go")
		f, err := os.Create(fp)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(f, goTest, name, s)

		if err := f.Close(); err != nil {
			return err
		}

		c := exec.Command("goimports", "-w", fp)
		c.Stdin = os.Stdin
		c.Stderr = os.Stderr
		c.Stdout = os.Stdout

		return c.Run()
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
	"fmt"
)

func %sExec(c *ast.Context) (*ast.Returned, error) {
	%#v
}
`

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
