package main

import rf "readfile"

type Dicebag struct {
	blue  int
	red   int
	green int
}

func partOne(lines []string) {
	dicebag := Dicebag{red: 12, green: 13, blue: 14}
	sum := 0
	for _, line := range lines {

	}
}

func main() {
	test1 := "test1.txt"
	lines := rf.ReadFile(test1)
	partOne(lines)
}
