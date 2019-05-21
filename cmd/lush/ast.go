package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gobuffalo/lush"
)

func printAST(args []string) {
	for _, a := range args {
		script, err := lush.ParseFile(a)
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.MarshalIndent(script, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}
}
