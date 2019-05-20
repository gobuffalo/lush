package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/lush/ast/internal/quick"
)

func main() {
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
