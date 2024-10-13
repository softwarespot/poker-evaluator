package poker

type Hand struct {
	kind  kind
	cards []card
}

// New creates a new hand; otherwise, returns an error, if the cards are invalid
func New(s string) (Hand, error) {
	cards, err := newCards(s)
	if err != nil {
		return Hand{
			kind:  kindUnknown,
			cards: nil,
		}, err
	}

	return Hand{
		kind:  getKind(cards),
		cards: cards,
	}, nil
}

// Compare returns 0 if both hands are a tie, -1 if the first hand won or 1, if the second hand won
func (h Hand) Compare(h2 Hand) Winner {
	if h.kind > h2.kind {
		return Winner1
	}
	if h.kind < h2.kind {
		return Winner2
	}

	// Same "kind"
	for i, c1 := range h.cards {
		c2 := h2.cards[i]
		if c1.ranking > c2.ranking {
			return Winner1
		}
		if c1.ranking < c2.ranking {
			return Winner2
		}
	}
	return WinnerTie
}
