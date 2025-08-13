package poker

type Winner int

const (
	WinnerTie Winner = 0
	Winner1   Winner = -1
	Winner2   Winner = 1
)

// String returns the winner as a string
func (w Winner) String() string {
	switch w {
	case WinnerTie:
		return "Tie"
	case Winner1:
		return "Hand 1"
	case Winner2:
		return "Hand 2"
	default:
		// NOTE: This shouldn't happen, due to the validation in "New()"
		return "Unknown"
	}
}
