package main

import (
	"log"

	"github.com/gobuffalo/lush/cmd/lush/commands"
)

func format(args []string) {
	r := commands.NewFmter(args)
	err := r.Exec()
	if err != nil {
		log.Fatal(err)
	}
}
