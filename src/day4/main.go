package main

import (
	util "adventofcode2023/util"
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	lines := util.ParseFile("input")
	maxLines := len(lines)
	scores := make(map[int]int)
	nbCards := make(map[int]int)
	for i := 1; i <= maxLines; i++ {
		nbCards[i] = 1
	}

	for lineIndex, line := range lines {
		gameIndex := lineIndex + 1
		cards := strings.Split(line, ":")
		splitCards := strings.Split(cards[1], "|")
		winningNumbers, handNumbers := strings.Fields(splitCards[0]), strings.Fields(splitCards[1])

		var wonNumbers []string
		for _, w := range winningNumbers {
			if slices.Contains(handNumbers, w) {
				wonNumbers = append(wonNumbers, w)
			}
		}

		if len(wonNumbers) == 0 {
			scores[gameIndex] = 0
		} else {
			scores[gameIndex] = int(math.Pow(2, float64(len(wonNumbers)-1)))

			// Part 2
			for j := 1; j <= nbCards[gameIndex]; j++ {
				for k := 1; k <= len(wonNumbers); k++ {
					if gameIndex+k <= maxLines {
						nbCards[gameIndex+k]++
					}
				}
			}
		}

	}

	total := 0
	for _, v := range scores {
		total += v
	}
	fmt.Println("Total Part 1", total)

	// fmt.Println(nbCards)
	totalNbCards := 0
	for _, v := range nbCards {
		totalNbCards += v
	}
	fmt.Println("Total Part 2", totalNbCards)
}
