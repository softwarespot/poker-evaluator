package poker

import (
	"fmt"
	"testing"

	testhelpers "github.com/softwarespot/poker-evaluator/test-helpers"
)

func Test_Hand_Compare(t *testing.T) {
	tests := []struct {
		cards1 string
		cards2 string
		want   string
	}{
		{"AC KC QC JC TC", "9D 8D 7D 6D 5D", "Hand 1"},
		{"JS JD JH JC 3S", "AS AD AH KS KD", "Hand 1"},
		{"KH JH 8H 4H 2H", "TS 9C 8D 7S 6C", "Hand 1"},
		{"QS QD 5C 5H 2D", "KS KD QC 9H 7S", "Hand 1"},
		{"AS JD 8C 4H 3S", "KD QC TH 6D 5S", "Hand 1"},
		{"AH AD AS KH KD", "AC AS AD KC KH", "Tie"},
		{"TH TS TD TC KH", "9C 9H 9S QD QH", "Hand 1"},
		{"5C 5D KH KS QC", "7H JD JS TC TD", "Hand 1"},
		{"3H 3S AC AD KH", "4C QD QS JH JD", "Hand 1"},
		{"6C KC QC JC TC", "7D TD TH TS TC", "Hand 2"},
		{"7C TH JH QH KH", "5S AS AD AC KD", "Hand 2"},
		{"AS KS QS JS TS", "KH QH JH TH AH", "Tie"},
		{"9C 8C 7C 6C 5C", "AD KD QD JD TD", "Hand 2"},
		{"2D AD AH AS KC", "3C QD QS JD JH", "Hand 1"},
		{"4S 4H KS TD JS", "4D AC AH KD QD", "Hand 2"},
		{"JD JH JS JC TH", "9C AC AH KS KD", "Hand 1"},
		{"3D TH JD QD KD", "3S TS JS QS KS", "Tie"},
		{"AD KD QD JD TD", "AS KS QS JS TS", "Tie"},
		{"2C 3C AC KC QC", "2H TH JH QH KH", "Hand 1"},
		{"8S AC AH AS KC", "8D AD AH AS QC", "Hand 1"},
		{"5H TH JH QH KH", "6S JS JD QS KD", "Hand 1"},
		{"9S TS JS QS KS", "8S TS JS QS KS", "Hand 1"},
		{"AS AH AC AD JC", "KS KD KH QD JD", "Hand 1"},
		{"2S TH JH QH KH", "3S TS JS QS KS", "Hand 2"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s vs %s = %s", tt.cards1, tt.cards2, tt.want), func(t *testing.T) {
			h1, err := New(tt.cards1)
			testhelpers.AssertNoError(t, err)

			h2, err := New(tt.cards2)
			testhelpers.AssertNoError(t, err)

			// Uncomment for better debugging
			// fmt.Printf("%+v vs %+v = %+v", h1, h1, h1.Compare(h2).String())
			testhelpers.AssertEqual(t, h1.Compare(h2).String(), tt.want)
		})
	}
}

func Test_Hand_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cards   string
		wantErr bool
	}{
		{name: "invalid rank", cards: "9C 9S PD 9H JH", wantErr: true},
		{name: "invalid suit", cards: "9C 9S 9D 9H JP", wantErr: true},
		{name: "invalid cards length (too many)", cards: "9C 9S 9D 9H JH 9C", wantErr: true},
		{name: "invalid cards length (too little)", cards: "9C 9S 9D 9H", wantErr: true},
		{name: "valid cards", cards: "9C 9S 9D 9H JH", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.cards)
			if tt.wantErr {
				testhelpers.AssertError(t, err)
			} else {
				testhelpers.AssertNoError(t, err)
			}
		})
	}
}
