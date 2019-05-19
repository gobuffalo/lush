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
	write(id, err)

	an, err := ast.NewAssign(id, num)
	write(an, err)

	acc, err := ast.NewAccess(id, 1)
	write(acc, err)
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

	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
