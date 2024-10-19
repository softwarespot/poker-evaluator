package poker

import (
	"fmt"
	"sort"
	"strings"
)

var (
	cardRanks = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	cardSuits = []rune{'C', 'D', 'H', 'S'}
)

type card struct {
	// Rank as a numerical value e.g. 2 is 0, K is 11, A is 12 etc.
	ranking int
	suit    rune
}

func newCards(s string) ([]card, error) {
	s = strings.ToUpper(strings.ReplaceAll(s, " ", ""))
	if len(s) != 10 {
		return nil, fmt.Errorf("invalid number of cards, got a length of %d for the cards %q", len(s), s)
	}

	var (
		chars = []rune(s)
		cards []card
	)
	for i := 0; i < len(chars); i += 2 {
		rank := chars[i]
		ranking := getRanking(rank)
		if ranking == -1 {
			return nil, fmt.Errorf("invalid card rank of %q was provided for the card %q in the cards %q", rank, chars[i:i+2], s)
		}

		suit := chars[i+1]
		if !isSuit(suit) {
			return nil, fmt.Errorf("invalid card suit of %q was provided for the card %q in the cards %q", suit, chars[i:i+2], s)
		}

		cards = append(cards, card{
			ranking: ranking,
			suit:    suit,
		})
	}

	// Sort first by the ranking and then the suit
	sort.SliceStable(cards, func(i, j int) bool {
		if cards[i].ranking == cards[j].ranking {
			return cards[i].suit < cards[j].suit
		}
		return cards[i].ranking > cards[j].ranking
	})
	return cards, nil
}

func getRanking(rank rune) int {
	for i, cardRank := range cardRanks {
		if cardRank == rank {
			return i
		}
	}
	return -1
}

func isSuit(suit rune) bool {
	for _, cardSuit := range cardSuits {
		if cardSuit == suit {
			return true
		}
	}
	return false
}

func getKind(cards []card) kind {
	uniqueRanksCount, maxCount := getUniqueRanksCountAndMaxCount(cards)
	switch uniqueRanksCount {
	case 2:
		if maxCount == 4 {
			return kindFourOfAKind
		}
		return kindFullHouse
	case 3:
		if maxCount == 3 {
			return kindThreeOfAKind
		}
		return kindTwoPair
	case 4:
		return kindOnePair
	default:
		if isRoyalFlush(cards) {
			return kindRoyalFlush
		}
		if isStraightFlush(cards) {
			return kindStraightFlush
		}
		if isFlush(cards) {
			return kindFlush
		}
		if isStraight(cards) {
			return kindStraight
		}
		return kindHigh
	}
}

func getUniqueRanksCountAndMaxCount(cards []card) (int, int) {
	var (
		countsByRank = map[int]int{}
		maxCount     int
	)
	for _, card := range cards {
		countsByRank[card.ranking]++
		if countsByRank[card.ranking] > maxCount {
			maxCount = countsByRank[card.ranking]
		}
	}
	return len(countsByRank), maxCount
}

func isRoyalFlush(cards []card) bool {
	// If a flush and the first and last cards are "A" and "10" respectively, then it's a royal flush.
	// This is due to ordering by rank and suit
	return isFlush(cards) && cards[0].ranking == getRanking('A') && cards[4].ranking == getRanking('T')
}

func isStraightFlush(cards []card) bool {
	return isFlush(cards) && isStraight(cards)
}

func isFlush(cards []card) bool {
	// If the first and last cards are the same suit, then it's a flush.
	// This is due to ordering by rank and suit
	return cards[0].suit == cards[4].suit
}

func isStraight(cards []card) bool {
	// Compare the previous card ranking minus 1 with the current card ranking.
	// If the same for all, then it's a straight.
	// This is due to ordering by rank and suit
	for i := 1; i < len(cards); i++ {
		currCard := cards[i]
		prevCard := cards[i-1]
		if currCard.ranking != prevCard.ranking-1 {
			return false
		}
	}
	return true
}
