package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gobuffalo/lush/ast"
)

func main() {
	s, err := ast.NewString([]byte(`"hi"`))
	write(s, err)

	arr, err := ast.NewArray([]interface{}{1, 2, 3})
	write(arr, err)

	id, err := ast.NewIdent([]byte("x"))
	write(id, err)

	num, err := ast.NewInteger(1)
	write(num, err)

	an, err := ast.NewAssign(id, num)
	write(an, err)

	vr, err := ast.NewVar(id, num)
	write(vr, err)

	acc, err := ast.NewAccess(id, 1)
	write(acc, err)

	block, err := ast.NewBlock(ast.Statements{an})
	write(block, err)

	br := ast.Break{}
	write(br, nil)

	nl := ast.Nil{}
	write(nl, nil)

	cmt := ast.Comment{Value: "hello"}
	write(cmt, nil)

	write(ast.True, nil)

	ctn := ast.Continue{}
	write(ctn, nil)

	mp := ast.Map{
		Values: map[ast.Statement]interface{}{
			id: num,
		},
	}
	write(mp, nil)
}

func write(s interface{}, err error) {
	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("%T", s)

	name = strings.TrimPrefix(name, "*")
	name = strings.TrimPrefix(name, "ast.")
	os.MkdirAll("ast/testdata", 0755)
	f, err := os.Create(fmt.Sprintf("ast/testdata/%s.json", name))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var b []byte
	if a, ok := s.(ast.ASTMarshaler); ok {
		b, err = a.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
	}

	if b == nil {
		b, err = json.MarshalIndent(s, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = f.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
