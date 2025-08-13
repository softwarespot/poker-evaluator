package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/softwarespot/poker-evaluator/internal/version"
)

func cmdVersion(asJSON bool) error {
	if !asJSON {
		fmt.Printf("Version:\t%s\n", version.Version)
		fmt.Printf("Build Time:\t%s\n", version.Time)
		fmt.Printf("Build User:\t%s\n", version.User)
		fmt.Printf("Go Version:\t%s\n", version.GoVersion)
		return nil
	}

	// This could be a struct, but it would be a temporary struct in that case.
	// Therefore, a map is honestly enough for this
	out := map[string]string{
		"version":   version.Version,
		"buildTime": version.Time,
		"buildUser": version.User,
		"goVersion": version.GoVersion,
	}
	return json.NewEncoder(os.Stdout).Encode(out)
}
