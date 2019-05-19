package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pig "github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
)

func main() {
	args := os.Args[1:]
	for _, a := range args {
		script, err := pig.ParseFile(a)
		if err != nil {
			log.Fatal(err)
		}
		c := ast.NewContext(context.Background(), os.Stdout)

		// fmt
		format(a, script)
		// fmt.Printf("%+v", script)
		// fmt

		res, err := script.Exec(c)
		if err != nil {
			log.Fatal(err)
		}

		if res == nil {
			return
		}

		if ri, ok := res.Value.([]interface{}); ok {
			for _, i := range ri {
				fmt.Println(i)
			}
			continue
		}
		fmt.Println(res)
	}
}

func format(a string, script ast.Script) {
	f, err := os.Create(a)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(script.String())
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
