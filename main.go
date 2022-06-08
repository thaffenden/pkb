package main

import (
	"log"

	"github.com/thaffenden/notes/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
