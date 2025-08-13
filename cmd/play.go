package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/softwarespot/poker-evaluator/internal/poker"
)

func cmdPlay(args []string, asJSON bool) error {
	cards1, cards2, err := getPlayersCards(args)
	if err != nil {
		return err
	}

	h1, err := poker.New(cards1)
	if err != nil {
		return err
	}

	h2, err := poker.New(cards2)
	if err != nil {
		return err
	}

	res := h1.Compare(h2)
	if !asJSON {
		fmt.Println(res)
		return nil
	}

	// This could be a struct, but it would be a temporary struct in that case.
	// Therefore, a map is honestly enough for this
	out := map[string]string{
		"hand1":  cards1,
		"hand2":  cards2,
		"result": res.String(),
	}
	return json.NewEncoder(os.Stdout).Encode(out)
}

func getPlayersCards(args []string) (string, string, error) {
	var withoutFlags []string
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			withoutFlags = append(withoutFlags, arg)
		}
	}
	if len(withoutFlags) != 2 {
		return "", "", errors.New("invalid number of arguments, expected at least 2 with the first hand and second hand")
	}
	return withoutFlags[0], withoutFlags[1], nil
}
