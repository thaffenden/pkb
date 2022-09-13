// Package main is the entrypoint for the CLI.
package main

import (
	"log"

	"github.com/thaffenden/pkb/cmd"
	"github.com/thaffenden/pkb/internal/config"
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
