package main

import (
	util "adventofcode2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var calibrationValues []int

func main() {
	lines := util.ParseFile("input2")
	digitReplacer := strings.NewReplacer(
		"oneight", "18",
		"twone", "21",
		"threeight", "38",
		"fiveight", "58",
		"sevenine", "79",
		"eightwo", "82",
		"eighthree", "83",
		"nineight", "98",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	letterRegex := regexp.MustCompile("[a-zA-Z]")

	for _, line := range lines {
		digits := letterRegex.ReplaceAllString(digitReplacer.Replace(line), "")
		calibrationString := digits[:1] + digits[len(digits)-1:]
		calibration, _ := strconv.ParseInt(calibrationString, 10, 0)
		calibrationValues = append(calibrationValues, int(calibration))
	}

	total := util.SumSliceInt(calibrationValues)
	fmt.Println("Total Part 2", total)
}
