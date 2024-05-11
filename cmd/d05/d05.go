package main

import (
	"fmt"
	rf "readfile"
	"strconv"
	"strings"
)

type TransformMap struct {
	name        string
	destination int
	source      int
	size        int
}

type InputMap struct {
	low   int
	size  int
	layer int
}

func linesToTransformMaps(lines []string) [][]TransformMap {
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
			size:        n[2],
		}
		transformMaps[tmLayer] = append(transformMaps[tmLayer], tmap)
	}

	return transformMaps
}

func sourceToDestination(source int, tmaps []TransformMap) int {
	destination := source

	for _, tmap := range tmaps {
		if source >= tmap.source && source < tmap.source+tmap.size {
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
	fmt.Println("Minimum destination location part one: ", mindest)
}

func sourceToDestinationRanges(sourceMap InputMap, inputMaps []InputMap, tmaps []TransformMap, layer int) (InputMap, []InputMap) {
	fmt.Println(layer, tmaps[0].name)
	destinationMap := sourceMap
	for _, tmap := range tmaps {
		fmt.Println("checking: ", sourceMap.low, sourceMap.low+sourceMap.size, "->", tmap.destination, tmap.destination+tmap.size)
		// if the sourceMap is entirely in the destinationMap
		//     <=======>
		// <===============>
		if sourceMap.low >= tmap.source && sourceMap.low+sourceMap.size < tmap.source+tmap.size {
			fmt.Println("SourceMap in target", tmap.source, tmap.source+tmap.size)
			destinationMap.low = sourceMap.low - tmap.source + tmap.destination
			destinationMap.size = sourceMap.size
		}
		// if source.low is before start and there is an overlap in the range
		// of the sourcemap in the destination let's make tha new InputMap and ad it to the array
		// <=======>
		//     <===============>
		if sourceMap.low < tmap.source && sourceMap.low+sourceMap.size > tmap.source && sourceMap.low+sourceMap.size < tmap.source+tmap.size {
			var tinMap = InputMap{
				sourceMap.low,
				tmap.source - sourceMap.low,
				layer,
			}
			inputMaps = append(inputMaps, tinMap)
			destinationMap.low = tmap.destination
			destinationMap.size = tmap.source - sourceMap.low
			fmt.Println("new inputMap leftside: ", tinMap)
		}
		// if the top end of source range is after the end of the transformMap and the low is
		// within the destination range then let's make that a new InputMap as well
		//            <=======>
		// <===============>
		if sourceMap.low > tmap.source && sourceMap.low < tmap.source+tmap.size && sourceMap.low+sourceMap.size > tmap.source+tmap.size {
			var tinMap = InputMap{
				tmap.destination + tmap.size,
				sourceMap.low + sourceMap.size - tmap.destination + tmap.size,
				layer,
			}
			inputMaps = append(inputMaps, tinMap)
			destinationMap.low = sourceMap.low - tmap.source + tmap.destination
			destinationMap.size = tmap.source + tmap.size - sourceMap.low
			fmt.Println("new inputMap rightside: ", tinMap)
		}
		// If the end of the sourceMap is before the start of the transformMap, skip the the map
		// <=======>
		//            <===============>
		if tmap.source > sourceMap.low+sourceMap.size {
			continue
		}
		// if the start of the sourceMap is after the end of the transformMap, skip the map
		//                    <=======>
		// <===============>
		if sourceMap.low > tmap.source+tmap.size {
			continue
		}
	}
	return destinationMap, inputMaps
}

func partTwo(lines []string) {
	// Reads first line of input file, splits the string into an array of Ints
	seedstr := strings.Split(lines[0][7:], " ")
	var seedArray []int
	for _, seed := range seedstr {
		s, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println(err)
			continue
		}
		seedArray = append(seedArray, s)
	}
	// takes the array of ints and converts them to an InputMap
	// Inputmaps are a start location and a size of a range
	var inputMaps []InputMap
	for i := 0; i < len(seedArray); i += 2 {
		inpmap := InputMap{seedArray[i], seedArray[i+1] - 1, 0}
		inputMaps = append(inputMaps, inpmap)
	}
	// Then read the rest of the file and create TransformMaps
	// TransforMaps have a description, destination, origin then size of range
	var destMap, minMap InputMap
	minMap.low = 1 << 31
	minMap.size = 10
	tmapsArray := linesToTransformMaps(lines[2:])
	for i := 0; i < len(inputMaps); i++ {
		fmt.Println("Starting new range: ", inputMaps[i], " at layer : ", inputMaps[i].layer, "\n\n")
		destMap = inputMaps[i]
		for n := inputMaps[i].layer; n < len(tmapsArray); n++ {
			destMap, inputMaps = sourceToDestinationRanges(destMap, inputMaps, tmapsArray[n], n)
			fmt.Println("\ninputMaps:", inputMaps, "\n\n")
		}
	}
	if destMap.low < minMap.low {
		minMap = destMap
	}
	fmt.Println(minMap, destMap)
}

func main() {
	lines := rf.ReadFile("test.txt")
	// lines := rf.ReadFile("input.txt")
	partOne(lines)
	partTwo(lines)
}
