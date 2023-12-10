package main

import (
	util "adventofcode2023/util"
	"fmt"
	"strconv"
	"strings"
)

type Mapping struct {
	dest     int
	src      int
	mapRange int
}

func newMapping(line []string) *Mapping {
	dest, _ := strconv.Atoi(line[0])
	src, _ := strconv.Atoi(line[1])
	mapRange, _ := strconv.Atoi(line[2])
	return &Mapping{dest: dest, src: src, mapRange: mapRange}
}

func main() {
	lines := util.ParseFile("example")
	var seeds []int
	for _, seedString := range strings.Fields(strings.Split(lines[0], ":")[1]) {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	var mapTypes [][]*Mapping
	var currentMap []*Mapping
	for _, lineString := range lines[2:] {
		// If empty line, move on to the next type of map to be filled
		if lineString == "" {
			mapTypes = append(mapTypes, currentMap)
			currentMap = nil
			continue
		}
		if strings.Contains(lineString, ":") {
			continue
		}

		// Store all maps for current type
		line := strings.Fields(lineString)
		currentMap = append(currentMap, newMapping(line))
	}

	minLocation := computeAllNumberMappings(seeds[0], mapTypes, 0)
	for _, seed := range seeds {
		val := computeAllNumberMappings(seed, mapTypes, 0)
		if val < minLocation {
			minLocation = val
		}
	}
	fmt.Println("Minimum Part 1:", minLocation)

	// Part 2
	minLocation2 := computeAllNumberMappings(seeds[0], mapTypes, 0)
	for i, seedEntry := range seeds {
		if i%2 == 1 {
			continue
		}
		nbIndexes := seeds[i+1]
		for seed := seedEntry; seed < seedEntry+nbIndexes-1; seed++ {
			val := computeAllNumberMappings(seed, mapTypes, 0)
			if val < minLocation2 {
				minLocation2 = val
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
	fmt.Println("Minimum Part 2:", minLocation2)
}

func computeAllNumberMappings(src int, mapTypes [][]*Mapping, index int) int {
	if index == len(mapTypes) {
		return src
	}
	val := computeNumberMapping(src, mapTypes[index])
	fmt.Printf("%d ", val)
	return computeAllNumberMappings(val, mapTypes, index+1)
}

func computeNumberMapping(src int, mappings []*Mapping) int {
	for _, mapping := range mappings {
		if src >= mapping.src && src <= mapping.src+mapping.mapRange-1 {
			return mapping.dest + (src - mapping.src)
		}
	}
	return src
}
