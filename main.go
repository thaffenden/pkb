// Package main is the entrypoint for the CLI.
package main

import (
	"fmt"
	"os"

	"github.com/thaffenden/pkb/cmd"
)

func main() {
	code := 0
	defer func() {
		os.Exit(code)
	}()

	if err := cmd.Execute(); err != nil {
		code = 1
		fmt.Printf("error: %v\n", err)
	}
}
