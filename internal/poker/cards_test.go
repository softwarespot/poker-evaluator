package poker

import (
	"fmt"
	"testing"

	testhelpers "github.com/softwarespot/poker-evaluator/test-helpers"
)

func Test_getKind(t *testing.T) {
	tests := []struct {
		cards string
		want  kind
	}{
		{"AD KD QD JD TD", kindRoyalFlush},
		{"QH JH TH 9H 8H", kindStraightFlush},
		{"9C 9S 9D 9H JH", kindFourOfAKind},
		{"3C 3S 3D 6C 6H", kindFullHouse},
		{"KC TC 7C 6C 4C", kindFlush},
		{"7C 6S 5S 4H 3H", kindStraight},
		{"2D 2S 2C KH 6H", kindThreeOfAKind},
		{"JH JC 4C 4S 9H", kindTwoPair},
		{"4H 4S KS TD 5S", kindOnePair},
		{"KH JH 8C 7D 4S", kindHigh},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s = %d", tt.cards, tt.want), func(t *testing.T) {
			cards, err := newCards(tt.cards)
			testhelpers.AssertNoError(t, err)
			got := getKind(cards)
			testhelpers.AssertEqual(t, got, tt.want)
		})
	}
}

func Test_isRoyalFlush(t *testing.T) {
	tests := []struct {
		cards string
		want  bool
	}{
		{"AH KH JH QH TH", true},
		{"9H KH JH QH TH", false},
		{"AH KH 8S QH TH", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s = %t", tt.cards, tt.want), func(t *testing.T) {
			cards, err := newCards(tt.cards)
			testhelpers.AssertNoError(t, err)
			got := isRoyalFlush(cards)
			testhelpers.AssertEqual(t, got, tt.want)
		})
	}
}

func Test_isStraightFlush(t *testing.T) {
	tests := []struct {
		cards string
		want  bool
	}{
		{"AH KH JH QH TH", true}, // A royal flush can also be a straight flush
		{"9H KH JH QH TH", true},
		{"AC KD QS JH TH", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s = %t", tt.cards, tt.want), func(t *testing.T) {
			cards, err := newCards(tt.cards)
			testhelpers.AssertNoError(t, err)
			got := isStraightFlush(cards)
			testhelpers.AssertEqual(t, got, tt.want)
		})
	}
}

func Test_isFlush(t *testing.T) {
	tests := []struct {
		cards string
		want  bool
	}{
		{"AH KH JH QH TH", true}, // A royal flush can also be a flush
		{"9S 8S JS QS TS", true},
		{"AC KD QS JH TH", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s = %t", tt.cards, tt.want), func(t *testing.T) {
			cards, err := newCards(tt.cards)
			testhelpers.AssertNoError(t, err)
			got := isFlush(cards)
			testhelpers.AssertEqual(t, got, tt.want)
		})
	}
}

func Test_isStraight(t *testing.T) {
	tests := []struct {
		cards string
		want  bool
	}{
		{"AH KH JH QH TH", true}, // A royal flush can also be a straight
		{"9S KS JS QS TS", true},
		{"AC KD QS 2H TH", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s = %t", tt.cards, tt.want), func(t *testing.T) {
			cards, err := newCards(tt.cards)
			testhelpers.AssertNoError(t, err)
			got := isStraight(cards)
			testhelpers.AssertEqual(t, got, tt.want)
		})
	}
}
