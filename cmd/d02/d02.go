package main

import (
	"fmt"
	rf "readfile"
	"strconv"
	"strings"
)

func partOne(lines []string) {
	diceBag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0

	for n, line := range lines {
		gameDice := map[string]int {
			"red": 0, 
			"green": 0, 
			"blue": 0,
		}
		game := line[strings.Index(line, ":")+2:]
		sets := strings.Split(string(game), "; ")

		// fmt.Println(sets)
		for _, set := range sets {
			// fmt.Println(set)
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				idn := strings.Index(cube, " ")
				cubecount, _ := strconv.Atoi(cube[:idn])
				cubecolour := cube[idn + 1:]
				gameDice[cubecolour] += cubecount
				// fmt.Println("game", n+1, ":", gameDice)
			}
		}
		if gameDice["red"] <= diceBag["red"] && gameDice["green"] <= diceBag["green"] && gameDice["blue"] <= diceBag["blue"] {
			sum += n+1
		}
	}
	fmt.Println(sum)
}

func main() {
	test1 := "test1.txt"
	lines := rf.ReadFile(test1)
	partOne(lines)
}
