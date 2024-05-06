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

func linesToTransformMaps(lines []string, searchstr string) []TransformMap {
	start := slices.Index(lines, searchstr) + 1
	var tmap []TransformMap

	for _, line := range lines[start:] {
		if strings.TrimSpace(line) == "" {
			break
		}
		valuesStr := strings.Split(line, " ")
		n := make([]int, len(valuesStr))
		for i, value := range valuesStr {
			n[i], _ = strconv.Atoi(value)
		}
		tmp := TransformMap{
			name:        searchstr,
			destination: n[0],
			source:      n[1],
			maprange:    n[2],
		}
		tmap = append(tmap, tmp)
	}
	return tmap
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
	seedToSoilMaps := linesToTransformMaps(lines, "seed-to-soil map:")
	soilToFertilizerMaps := linesToTransformMaps(lines, "soil-to-fertilizer map:")
	fertilizerToWaterMaps := linesToTransformMaps(lines, "fertilizer-to-water map:")
	waterToLightMaps := linesToTransformMaps(lines, "water-to-light map:")
	lightToTemperatureMaps := linesToTransformMaps(lines, "light-to-temperature map:")
	temperatureToHumidityMaps := linesToTransformMaps(lines, "temperature-to-humidity map:")
	humidityToLocationMaps := linesToTransformMaps(lines, "humidity-to-location map:")

	var mindest, dest int
	mindest = 1 << 31
	for _, seed := range seeds {
		dest = sourceToDestination(seed, seedToSoilMaps)
		dest = sourceToDestination(dest, soilToFertilizerMaps)
		dest = sourceToDestination(dest, fertilizerToWaterMaps)
		dest = sourceToDestination(dest, waterToLightMaps)
		dest = sourceToDestination(dest, lightToTemperatureMaps)
		dest = sourceToDestination(dest, temperatureToHumidityMaps)
		dest = sourceToDestination(dest, humidityToLocationMaps)
		if dest < mindest {
			mindest = dest
		}
	}
	fmt.Println("Minimum location: ", mindest)
}

func main() {
	// lines := rf.ReadFile("test.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
}
