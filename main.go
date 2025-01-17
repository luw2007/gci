package main

import (
	"os"

	"github.com/luw2007/gci/cmd/gci"
)

var version = "0.9.0"

func main() {
	e := gci.NewExecutor(version)

	err := e.Execute()
	if err != nil {
		os.Exit(1)
	}
}
