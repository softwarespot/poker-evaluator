package cmd

import "fmt"

func cmdHelp() {
	helpText := `Usage: ./poker-evaluator [OPTIONS] HAND1 HAND2

Compare two poker hands and determine the winning hand (unless there is a tie).

Arguments:
  HAND1       The first hand to evaluate (e.g. "AH AC AD QC QD").
  HAND2       The second hand to evaluate (e.g. "QD QH QC AD AC").

Options:
  -h, --help      Show this help text and exit.
  -v, --version   Display the version of the application and exit.
  -j, --json      Output the result as JSON.

Examples:
  ./poker-evaluator "AH AC AD QC QD" "QD QH QC AD AC"
  ./poker-evaluator --json "AH AC AD QC QD" "QD QH QC AD AC"`
	fmt.Println(helpText)
}
