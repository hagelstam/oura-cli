package main

import (
	"os"

	"github.com/hagelstam/oura-cli/cmd/oura"
)

func main() {
	if err := oura.Execute(); err != nil {
		os.Exit(1)
	}
}
