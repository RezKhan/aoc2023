package main

import (
	"fmt"
	rf "readfile"
	"slices"
	"strconv"
	"strings"
)

type TransformMap struct {
	name        string
	destination int
	source      int
	maprange    int
}

func linesToTransformMaps (lines []string) [][]TransformMap {
	tname := ""
	tmLayer := 0
	var transformMaps [][]TransformMap
	var emptyMaps []TransformMap

	transformMaps = append(transformMaps, emptyMaps)
	for _, line := range lines {
		if strings.Contains(line, ":") {
			tname = line
			continue
		}
		
		if strings.TrimSpace(line) == "" {
			tmLayer++
			transformMaps = append(transformMaps, emptyMaps)
			continue
		}

		valuesStr := strings.Split(line, " ")
		n := make([]int, len(valuesStr))
		for i, value := range valuesStr {
			n[i], _ = strconv.Atoi(value)
		}
		tmap := TransformMap{
			name:        tname,
			destination: n[0],
			source:      n[1],
			maprange:    n[2],
		}
		transformMaps[tmLayer] = append(transformMaps[tmLayer], tmap)
	}

	return transformMaps
}

func sourceToDestination(source int, tmaps []TransformMap) int {
	destination := source

	for _, tmap := range tmaps {
		if source >= tmap.source && source < tmap.source+tmap.maprange {
			destination = source - tmap.source + tmap.destination
			return destination
		} else {
			continue
		}
	}

	return destination
}

func partOne(lines []string) {
	seedstr := strings.Split(lines[0][7:], " ")
	var seeds []int

	for _, seed := range seedstr {
		s, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println(err)
			continue
		}
		seeds = append(seeds, s)
	}

	var mindest, dest int
	mindest = 1 << 31
	tmapsarr := linesToTransformMaps(lines[2:])
	for _, seed := range seeds {
		dest = seed
		for _, tmaps := range tmapsarr {
			dest = sourceToDestination(dest, tmaps)
		}
		if dest < mindest {
			mindest = dest
		}
	}
	fmt.Println("Minimum destination location: ", mindest)
}

func partTwo (lines []string) {
	seedstr := strings.Split(lines[0][7:], " ")
	var seedsarr []int

	for _, seed := range seedstr {
		s, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println(err)
			continue
		}
		seedsarr = append(seedsarr, s)
	}
	tmapsarr := linesToTransformMaps(lines[2:])
	for _, tmaps := range tmapsarr {
		for _, tmap := range tmaps {
			fmt.Println(tmap)
		}
	}

}

func main() {
	lines := rf.ReadFile("test.txt")
	// lines := rf.ReadFile("input.txt")
	partOne(lines)
	partTwo(lines)
}
