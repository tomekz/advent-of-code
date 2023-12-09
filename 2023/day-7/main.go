package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type HandType string

const (
	FiveOfAKind  HandType = "Five of a kind"
	FourOfAKind  HandType = "Four of a kind"
	FullHouse    HandType = "Full house"
	ThreeOfAKind HandType = "Three of a kind"
	TwoPair      HandType = "Two pair"
	OnePair      HandType = "One pair"
	HighCard     HandType = "High card"
)

type Label string

const (
	A     Label = "A"
	K     Label = "K"
	Q     Label = "Q"
	J     Label = "J"
	T     Label = "T"
	Nine  Label = "9"
	Eight Label = "8"
	Seven Label = "7"
	Six   Label = "6"
	Five  Label = "5"
	Four  Label = "4"
	Three Label = "3"
	Two   Label = "2"
)

type Card struct {
	label    Label
	strength int
}

type Hand struct {
	cards []Card
	bid   int
}

var CamelCards = map[Label]int{
	Two:   2,
	Three: 3,
	Four:  4,
	Five:  5,
	Six:   6,
	Seven: 7,
	Eight: 8,
	Nine:  9,
	T:     10,
	J:     11,
	Q:     12,
	K:     13,
	A:     14,
}

var HandsTypes = map[HandType]int{
	FiveOfAKind:  7,
	FourOfAKind:  6,
	FullHouse:    5,
	ThreeOfAKind: 4,
	TwoPair:      3,
	OnePair:      2,
	HighCard:     1,
}

func (c *Card) Strength() int {
	return CamelCards[c.label]
}

// func (h *Hand) Type() HandType {
// 	cardCounts := make(map[Label]int)
// 	for _, card := range h.cards {
// 		cardCounts[card.label]++
// 	}
// }

func (h *Hand) Type() HandType {
	cardCounts := make(map[Label]int)
	for _, card := range h.cards {
		cardCounts[card.label]++
	}

	uniq := len(cardCounts)

	if uniq == 1 {
		return FiveOfAKind
	}

	if uniq == 5 {
		return HighCard
	}

	if uniq == 4 {
		return OnePair
	}

	if uniq == 3 {
		for _, count := range cardCounts {
			if count == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	}

	if uniq == 2 {
		for _, count := range cardCounts {
			if count == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	}

	return HighCard
}

func (h *Hand) Strength() int {
	var strength int
	for _, card := range h.cards {
		strength += card.Strength()
	}
	return strength
}

func (h *Hand) Less(h2 Hand) bool {
	if h.Type() == h2.Type() {
		return h.Strength() < h2.Strength()
	}
	return HandsTypes[h.Type()] < HandsTypes[h2.Type()]
}

func parseInput() []Hand {
	hands := []Hand{}
	contents, _ := os.ReadFile("input")
	lines := strings.Split(string(contents), "\n")[:1000]

	for _, line := range lines {
		s := strings.Split(line, " ")
		cards := s[0]
		bid := s[1]

		hand := Hand{}
		b, _ := strconv.Atoi(bid)
		hand.bid = b
		for _, card := range cards {
			c := Card{label: Label(card), strength: CamelCards[Label(card)]}
			hand.cards = append(hand.cards, c)
		}
		hands = append(hands, hand)
	}
	return hands
}

func main() {
	hands := parseInput()

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Less(hands[j])
	})
	total := 0
	for i, h := range hands {
		rank := i + 1
		total += rank * h.bid
	}

	fmt.Println(total)
}
