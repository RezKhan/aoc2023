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

func partOne(lines []string) {
	sum := 0
	for _, line := range lines {
		winningNumsStr := line[strings.Index(line, ":")+2:strings.Index(line, "|") - 1]
		heldNumsStr := line[strings.Index(line, "|") + 2:]
		winningNums := cardNumStrToInts(winningNumsStr)
		heldNums := cardNumStrToInts(heldNumsStr)
		fmt.Println(heldNums, "|", winningNums)
		
		t := 0
		for _, held := range heldNums {
			if slices.Index(winningNums, held) > -1 {
				if t < 1 {
					t = 1
				} else {
					t *= 2
				}
			}
		}
		sum += t
		fmt.Println(t, sum)
	}
	fmt.Println(sum)
}

func main() {
	// lines := rf.ReadFile("test.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
}
