package main

import (
	"fmt"
	rf "readfile"
	"strconv"
	"strings"
)

func partOne(lines []string) {
	bagDice := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0

	for n, line := range lines {
		
		gameSuccess := true
		game := line[strings.Index(line, ":")+2:]
		sets := strings.Split(string(game), "; ")

		// fmt.Println(sets)
		for _, set := range sets {
			gameDice := map[string]int {
				"red": 0, 
				"green": 0, 
				"blue": 0,
			}
			// fmt.Println(set)
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				idx := strings.Index(cube, " ")
				cubecount, _ := strconv.Atoi(cube[:idx])
				cubecolour := cube[idx + 1:]
				gameDice[cubecolour] += cubecount
				// fmt.Println("game", n+1, ":", gameDice)
			}
			if gameDice["red"] > bagDice["red"] || gameDice["green"] > bagDice["green"] || gameDice["blue"] > bagDice["blue"] {
				gameSuccess = false
			}
		}
		if gameSuccess == true {
			sum += n+1
		}
	}
	fmt.Println(sum)
}



func partTwo(lines []string) {
	sum := 0

	for _, line := range lines {
		
		game := line[strings.Index(line, ":")+2:]
		sets := strings.Split(string(game), "; ")
		// fmt.Println(sets)
		
		minDice := map[string]int {
			"red": 0, 
			"green": 0, 
			"blue": 0,
		}
		
		for _, set := range sets {
			gameDice := map[string]int {
				"red": 0, 
				"green": 0, 
				"blue": 0,
			}
			// fmt.Println(set)
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				idx := strings.Index(cube, " ")
				cubecount, _ := strconv.Atoi(cube[:idx])
				cubecolour := cube[idx + 1:]
				gameDice[cubecolour] += cubecount
				if gameDice[cubecolour] > minDice[cubecolour] {
					minDice[cubecolour] = gameDice[cubecolour]
				}
			}
		}
		sum += minDice["red"] * minDice["green"] * minDice["blue"]
	}
	fmt.Println(sum)
}

func main() {
	// test1 := "test1.txt"
	// lines := rf.ReadFile(test1)
	inputtxt := "input.txt"
	lines := rf.ReadFile(inputtxt)
	partOne(lines)
	partTwo(lines)
}
