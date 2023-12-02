package main

import (
	util "adventofcode2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var calibrationValues1 []int
var calibrationValues2 []int

func main() {
	lines := util.ParseFile("input")

	// Part 1

	// Method 1 - Find all digits
	// digitRegex := regexp.MustCompile("[0-9]")
	// for _, line := range lines {
	// 	digits := digitRegex.FindAllString(line, -1)
	// 	calibration, _ := strconv.ParseInt(digits[0]+digits[len(digits)-1], 10, 0)
	// 	calibrationValues = append(calibrationValues, int(calibration))
	// }

	// Method 2 - Remove all letters
	letterRegex := regexp.MustCompile("[a-zA-Z]")
	for _, line := range lines {
		digits := letterRegex.ReplaceAllString(line, "")
		calibrationString := digits[:1] + digits[len(digits)-1:]
		calibration, _ := strconv.ParseInt(calibrationString, 10, 0)
		calibrationValues1 = append(calibrationValues1, int(calibration))
	}

	total := util.SumSliceInt(calibrationValues1)
	fmt.Println("Total Part 1", total)

	// Part 2
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
	for _, line := range lines {
		digits := letterRegex.ReplaceAllString(digitReplacer.Replace(line), "")
		calibrationString := digits[:1] + digits[len(digits)-1:]
		calibration, _ := strconv.ParseInt(calibrationString, 10, 0)
		calibrationValues2 = append(calibrationValues2, int(calibration))
	}

	total2 := util.SumSliceInt(calibrationValues2)
	fmt.Println("Total Part 2", total2)
}
