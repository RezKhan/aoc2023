package main

import (
	"fmt"
	rf "readfile"
	"slices"
	"strconv"
	"strings"
)

func cardNumStrToInts(str string) []int {
	var nums []int
	strs := strings.Split(str, " ")
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		nums = append(nums, n)
	}

	return nums
}

func assessCardForWins(winningNums []int, heldNums[]int) (winValue int, winCount int) {
	winValue = 0
	winCount = 0
	for _, held := range heldNums {
		if slices.Index(winningNums, held) > -1 {
			if winValue < 1 {
				winValue = 1
			} else {
				winValue *= 2
			}
			winCount++
		}
	}
	return winValue, winCount
}


func partOne(lines []string) {
	sum := 0
	for _, line := range lines {
		winningNumsStr := line[strings.Index(line, ":")+2 : strings.Index(line, "|")-1]
		heldNumsStr := line[strings.Index(line, "|")+2:]
		winningNums := cardNumStrToInts(winningNumsStr)
		heldNums := cardNumStrToInts(heldNumsStr)
		winValue, _ := assessCardForWins(winningNums, heldNums)

		sum += winValue
	}
	fmt.Println(sum)
}

func partTwo(lines []string) {
	var copies []int
	for range lines {
		copies = append(copies, 0)
	}
	for i, line := range lines {
		winningNumsStr := line[strings.Index(line, ":")+2 : strings.Index(line, "|")-1]
		heldNumsStr := line[strings.Index(line, "|")+2:]
		winningNums := cardNumStrToInts(winningNumsStr)
		heldNums := cardNumStrToInts(heldNumsStr)
		for j := 0; j <= copies[i]; j++ {
			_, winCount := assessCardForWins(winningNums, heldNums)
			for n:= 1; n <= winCount; n++ {
				copies[i+n]++
			}
		}
	}
	totalCards := 0
	for i := range copies {
		totalCards += 1 + copies[i]
	}

	fmt.Println(totalCards)
}

func main() {
	// lines := rf.ReadFile("test.txt")
	lines := rf.ReadFile("input.txt")
	// partOne(lines)
	partTwo(lines)
}
