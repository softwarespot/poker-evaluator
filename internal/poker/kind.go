package poker

// "kind" is used, because "type" is a keyword and "handType" doesn't sound good as "poker.handType"
type kind int

// The order is from the lowest hand to the highest hand, as comparison of the kind,
// will be semantically be correct i.e. h1.kind > h2.kind
const (
	kindUnknown kind = iota

	// Example: KH JH 8C 7D 4S
	kindHigh

	// Example: 4H 4S KS TD 5S
	kindOnePair

	// Example: JH JC 4C 4S 9H
	kindTwoPair

	// Example: 2D 2S 2C KH 6H
	kindThreeOfAKind

	// Example: 7C 6S 5S 4H 3H
	kindStraight

	// Example: KC TC 7C 6C 4C
	kindFlush

	// Example: 3C 3S 3D 6C 6H
	kindFullHouse

	// Example: 9C 9S 9D 9H JH
	kindFourOfAKind

	// Example: QH JH TH 9H 8H
	kindStraightFlush

	// Example: AD KD QD JD TD
	kindRoyalFlush
)
