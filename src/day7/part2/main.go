package main

import (
	util "adventofcode2023/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	FIVE_KIND  = 6
	FOUR_KIND  = 5
	FULL_HOUSE = 4
	THREE_KIND = 3
	TWO_PAIR   = 2
	ONE_PAIR   = 1
	HIGH_CARD  = 0
)

var cardStrength = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type Hand struct {
	cards string
	bid   int
}

func newHand(line string) *Hand {
	lineSplit := strings.Fields(line)
	bid, _ := strconv.Atoi(lineSplit[1])
	return &Hand{cards: lineSplit[0], bid: bid}
}

func main() {
	lines := util.ParseFile("../input")
	handTypes := make([][]*Hand, 7)
	for _, line := range lines {
		hand := newHand(line)
		handType := getHandType(hand.cards)
		handTypes[handType] = append(handTypes[handType], hand)

		// Sort hands for each hand type by strength
		for _, hands := range handTypes {
			sort.Slice(hands, func(i, j int) bool {
				if cardStrength[string(hands[i].cards[0])] == cardStrength[string(hands[j].cards[0])] {
					if cardStrength[string(hands[i].cards[1])] == cardStrength[string(hands[j].cards[1])] {
						if cardStrength[string(hands[i].cards[2])] == cardStrength[string(hands[j].cards[2])] {
							if cardStrength[string(hands[i].cards[3])] == cardStrength[string(hands[j].cards[3])] {
								return cardStrength[string(hands[i].cards[4])] < cardStrength[string(hands[j].cards[4])]
							}
							return cardStrength[string(hands[i].cards[3])] < cardStrength[string(hands[j].cards[3])]
						}
						return cardStrength[string(hands[i].cards[2])] < cardStrength[string(hands[j].cards[2])]
					}
					return cardStrength[string(hands[i].cards[1])] < cardStrength[string(hands[j].cards[1])]
				}
				return cardStrength[string(hands[i].cards[0])] < cardStrength[string(hands[j].cards[0])]
			})
		}
	}

	var rankedHands []*Hand
	for _, handType := range handTypes {
		rankedHands = append(rankedHands, handType...)
	}

	total := 0
	for i, h := range rankedHands {
		// fmt.Println(h.cards, h.bid, i+1)
		total += h.bid * (i + 1)
	}
	fmt.Println("Total Part 2", total)
}

func getHandType(cards string) int {
	frequency := make(map[string]int)
	for _, card := range cards {
		frequency[string(card)]++
	}

	// Get max card other than J
	if frequency["J"] != 0 {
		maxCard := ""
		maxNbCards := 0
		for card, f := range frequency {
			if card != "J" && f > maxNbCards {
				maxCard = card
				maxNbCards = f
			}
		}
		frequency[maxCard] += frequency["J"]
		delete(frequency, "J")
	}

	// Compute hand type
	if len(frequency) == 1 {
		return FIVE_KIND
	}
	if len(frequency) == 2 {
		isFourKind := false
		for _, f := range frequency {
			if f == 4 {
				isFourKind = true
			}
		}
		if isFourKind {
			return FOUR_KIND
		} else {
			return FULL_HOUSE
		}
	}
	if len(frequency) == 3 {
		isThreeKind := false
		for _, f := range frequency {
			if f == 3 {
				isThreeKind = true
			}
		}
		if isThreeKind {
			return THREE_KIND
		} else {
			return TWO_PAIR
		}
	}
	if len(frequency) == 4 {
		return ONE_PAIR
	}
	return HIGH_CARD
}
