package main

import (
	"os"

	"github.com/softwarespot/poker-evaluator/cmd"
	"github.com/softwarespot/poker-evaluator/internal/helpers"
)

func main() {
	// Remove the executable name
	if err := cmd.Execute(os.Args[1:]); err != nil {
		helpers.FatalExit(err, 1)
	}
}
