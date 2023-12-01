package main

import (
	util "adventofcode2023/util"
	"fmt"
	"regexp"
	"strconv"
)

var calibrationValues []int

func main() {
	lines := util.ParseFile("input1")

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
		calibrationValues = append(calibrationValues, int(calibration))
	}

	total := util.SumSliceInt(calibrationValues)
	fmt.Println("Total Part 1", total)
}
