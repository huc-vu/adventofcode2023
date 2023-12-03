package main

import (
	util "adventofcode2023/util"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

var total = 0
var totalPower = 0

func main() {
	lines := util.ParseFile("input")
	games := make(map[int]string)
	for i, e := range lines {
		game := strings.Split(e, ":")
		games[i+1] = game[1]
	}
	for gameNumber, game := range games {
		redRegex := regexp.MustCompile(`(\d+) red`)
		redDices := countColors(*redRegex, game)

		greenRegex := regexp.MustCompile(`(\d+) green`)
		greenDices := countColors(*greenRegex, game)

		blueRegex := regexp.MustCompile(`(\d+) blue`)
		blueDices := countColors(*blueRegex, game)

		if slices.Max(redDices) <= MAX_RED && slices.Max(greenDices) <= MAX_GREEN && slices.Max(blueDices) <= MAX_BLUE {
			total += gameNumber
		}

		// Part 2
		minRed := slices.Max(redDices)
		minGreen := slices.Max(greenDices)
		minBlue := slices.Max(blueDices)
		power := minRed * minGreen * minBlue
		totalPower += power
	}
	fmt.Println("Total Part 1", total)
	fmt.Println("Total Part 2", totalPower)

}

func countColors(regex regexp.Regexp, game string) []int {
	var colorDices []int
	for _, occ := range regex.FindAllStringSubmatch(game, -1) {
		conv, _ := strconv.ParseInt(occ[1], 10, 0)
		colorDices = append(colorDices, int(conv))
	}
	return colorDices
}
