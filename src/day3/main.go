package main

import (
	util "adventofcode2023/util"
	"fmt"
	"regexp"
	"strconv"
)

var validNumbers []int

type Character struct {
	value string
	row   int
	col   int
}
type Engine struct {
	data  [][]string
	nbRow int
	nbCol int
}

func NewEngine(lines []string) *Engine {
	newData := make([][]string, len(lines))
	for rowIndex, line := range lines {
		newData[rowIndex] = make([]string, len(line))
		for colIndex, val := range line {
			newData[rowIndex][colIndex] = string(val)
		}
	}
	return &Engine{
		data: newData, nbRow: len(lines), nbCol: len(lines[0]),
	}
}

func main() {
	lines := util.ParseFile("input")
	engine := NewEngine(lines)
	gears := make(map[string][]int)

	digitRegex := regexp.MustCompile(`\d`)
	symbolRegex := regexp.MustCompile(`\+|\*|\=|\&|\#|\$|\-|\@|\/|\%`)
	allSymbolsRegex := regexp.MustCompile(`\+|\*|\=|\&|\#|\$|\-|\@|\/|\%|\.`)

	var currentNumber string
	isNumberValid := false
	isGear := false
	var gearPosition Character

	for rowIndex, row := range engine.data {
		for colIndex, val := range row {
			if digitRegex.MatchString(val) {
				currentNumber += string(val)
				surroundingCharacters := getSurroundingCharacters(engine, rowIndex, colIndex)
				// fmt.Println(val, surroundingCharacters)
				for _, c := range surroundingCharacters {
					isNumberValid = isNumberValid || symbolRegex.MatchString(c.value)
					if c.value == "*" {
						isGear = true
						gearPosition = c
					}
				}

				// Check character after number
				if surroundingCharacters["Right"].value != "" && allSymbolsRegex.MatchString(surroundingCharacters["Right"].value) || colIndex == engine.nbCol-1 {
					// fmt.Printf("current Number %s %s", currentNumber, surroundingCharacters["Right"])
					// fmt.Println("isNumberValid ", isNumberValid)
					if isNumberValid {
						parsedNumber, _ := strconv.Atoi(currentNumber)
						validNumbers = append(validNumbers, int(parsedNumber))
						if isGear {
							key := strconv.Itoa(gearPosition.row) + "-" + strconv.Itoa(gearPosition.col)
							gears[key] = append(gears[key], int(parsedNumber))
						}
					}
					isNumberValid = false
					isGear = false
					currentNumber = ""
				}
			}
		}
	}
	fmt.Println("Total Part 1", util.SumSliceInt(validNumbers))

	// Part 2
	gearRatios := 0
	for _, gear := range gears {
		if len(gear) == 2 {
			gearRatios += gear[0] * gear[1]
		}
	}
	fmt.Println("Total Part 2", gearRatios)
}

func getSurroundingCharacters(engine *Engine, rowIndex int, colIndex int) map[string]Character {
	return map[string]Character{
		"UpLeft":    getUpLeft(engine, rowIndex, colIndex),
		"Up":        getUp(engine, rowIndex, colIndex),
		"UpRight":   getUpRight(engine, rowIndex, colIndex),
		"Left":      getLeft(engine, rowIndex, colIndex),
		"Right":     getRight(engine, rowIndex, colIndex),
		"DownLeft":  getDownLeft(engine, rowIndex, colIndex),
		"Down":      getDown(engine, rowIndex, colIndex),
		"DownRight": getDownRight(engine, rowIndex, colIndex),
	}
}
func getUpLeft(engine *Engine, rowIndex int, colIndex int) Character {
	if rowIndex-1 < 0 || colIndex-1 < 0 {
		return Character{"", rowIndex - 1, colIndex - 1}
	}
	return Character{engine.data[rowIndex-1][colIndex-1], rowIndex - 1, colIndex - 1}
}
func getUp(engine *Engine, rowIndex int, colIndex int) Character {
	if rowIndex-1 < 0 {
		return Character{"", rowIndex - 1, colIndex}
	}
	return Character{engine.data[rowIndex-1][colIndex], rowIndex - 1, colIndex}
}
func getUpRight(engine *Engine, rowIndex int, colIndex int) Character {
	if rowIndex-1 < 0 || colIndex+1 > engine.nbCol-1 {
		return Character{"", rowIndex - 1, colIndex + 1}
	}
	return Character{engine.data[rowIndex-1][colIndex+1], rowIndex - 1, colIndex + 1}
}
func getLeft(engine *Engine, rowIndex int, colIndex int) Character {
	if colIndex-1 < 0 {
		return Character{"", rowIndex, colIndex - 1}
	}
	return Character{engine.data[rowIndex][colIndex-1], rowIndex, colIndex - 1}
}
func getRight(engine *Engine, rowIndex int, colIndex int) Character {
	if colIndex+1 > engine.nbCol-1 {
		return Character{"", rowIndex, colIndex + 1}
	}
	return Character{engine.data[rowIndex][colIndex+1], rowIndex, colIndex + 1}
}
func getDownLeft(engine *Engine, rowIndex int, colIndex int) Character {
	if rowIndex+1 > engine.nbRow-1 || colIndex-1 < 0 {
		return Character{"", rowIndex + 1, colIndex - 1}
	}
	return Character{engine.data[rowIndex+1][colIndex-1], rowIndex + 1, colIndex - 1}
}
func getDown(engine *Engine, rowIndex int, colIndex int) Character {
	if rowIndex+1 > engine.nbRow-1 {
		return Character{"", rowIndex + 1, colIndex}
	}
	return Character{engine.data[rowIndex+1][colIndex], rowIndex + 1, colIndex}
}
func getDownRight(engine *Engine, rowIndex int, colIndex int) Character {
	if rowIndex+1 > engine.nbRow-1 || colIndex+1 > engine.nbCol-1 {
		return Character{"", rowIndex + 1, colIndex + 1}
	}
	return Character{engine.data[rowIndex+1][colIndex+1], rowIndex + 1, colIndex + 1}
}
