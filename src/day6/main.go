package main

import (
	util "adventofcode2023/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := util.ParseFile("input")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])

	options := make([]int, len(times))
	for i, timeString := range times {
		time, _ := strconv.Atoi(timeString)
		distance, _ := strconv.Atoi(distances[i])
		options[i] = computeOptions(time, distance)
	}
	totalOptions := 1
	for _, v := range options {
		totalOptions *= v
	}
	fmt.Println("Total Part 1", totalOptions)

	// Part 2
	times2, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	distances2, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))
	fmt.Println("Total Part 2", computeOptions(times2, distances2))
}

func computeOptions(time int, distance int) int {
	nbOptions := 0
	for ms := 0; ms < time; ms++ {
		if ms*(time-ms) > distance {
			nbOptions++
		}
	}
	return nbOptions
}
