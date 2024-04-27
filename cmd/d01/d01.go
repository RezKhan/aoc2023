package main

import (
	"fmt"
	rf "readfile"
	"strconv"
	"strings"
)

func partOneNumber(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

func partOne(lines []string) {
	var sum int
	for _, line := range lines {
		var startNum int
		var endNum int
		lenline := len(line) - 1
		i := 0

		for startNum == 0 || endNum == 0 {
			if startNum == 0 {
				startNum = partOneNumber(string(line[i]))
			}
			if endNum == 0 {
				endNum = partOneNumber(string(line[lenline - i]))
			}
			i++
		}

		sum += (startNum * 10) + endNum
	}
	fmt.Println(sum)
}

func partTwoNumber(c string, substr string) int {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	n, err := strconv.Atoi(c)
	if err == nil {
		return n
	}

	for i := 0; i < len(numbers); i++ {
		if strings.Contains(substr, numbers[i]) {
			return i + 1
		}
	}

	return 0
}

func partTwo(lines []string) {
	var sum int
	for _, line := range lines {
		var startNum int
		var endNum int
		lenline := len(line) - 1
		i := 0

		// for i := 0; i < len(line); i++ {
		// 	fmt.Println(string(line[i]), line[:i+1])
		// 	fmt.Println(string(line[lenline-i]), line[lenline - i:])
		// }
		for startNum == 0 || endNum == 0 {
			if startNum == 0 {
				startNum = partTwoNumber(string(line[i]), line[:i+1])
			}
			if endNum == 0 {
				endNum = partTwoNumber(string(line[lenline-i]), line[lenline - i:])
			}
			i++
		}

		sum += (startNum * 10) + endNum
	}
	fmt.Println(sum)
}

func main() {
	// test1 := "test1.txt"
	// lines := rf.ReadFile(test1)
	// partOne(lines)
	// test2 := "test2.txt"
	// lines := rf.ReadFile(test2)
	// partTwo(lines)
	inputtxt := "input.txt"
	lines := rf.ReadFile(inputtxt)
	partOne(lines)
	partTwo(lines)

}
