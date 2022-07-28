package main

import (
	"log"

	"github.com/thaffenden/notes/cmd"
	"github.com/thaffenden/notes/internal/config"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(config); err != nil {
		log.Fatal(err)
	}
}
