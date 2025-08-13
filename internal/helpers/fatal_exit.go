package helpers

import (
	"fmt"
	"os"
)

func FatalExit(err error, code int) {
	fmt.Fprintf(os.Stderr, "Error:\n\t%s\n", err.Error())
	os.Exit(code)
}
