package main

import (
	"log"
	"os"

	"github.com/gobuffalo/lush/cmd/lush/commands"
)

func main() {
	args := os.Args[1:]
	if err := commands.Route(args); err != nil {
		log.Fatal(err)
	}
}
