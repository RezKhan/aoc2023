package main

import (
	"fmt"
	rf "readfile"
	"strconv"
	"unicode"
)

type Coords struct {
	y int
	x int
}
type PartNumber struct {
	value    int
	coords   Coords
	len      int
	adjacent bool
}

type SymbolLocation struct {
	value  string
	coords Coords
}

func partOne(lines []string) {
	var slocs []SymbolLocation
	var pnums []PartNumber
	sum := 0

	// rows := len(lines)
	cols := len(lines[0])

	// Scan lines for symbols, ignore periods
	for y, line := range lines {
		num := 0
		count := 0
		for x := 0; x < cols; x++ {
			n, err := strconv.Atoi(string(line[x]))
			if err == nil {
				num = (num * 10) + n
				count++

				if x >= cols -1 || !unicode.IsDigit(rune(line[x+1])){
					var pnum PartNumber
					pnum.value = num
					pnum.len = count - 1 
					pnum.coords.x = x - count + 1
					pnum.coords.y = y
					pnums = append(pnums, pnum)
					num = 0
					count = 0
				}
			}

			if string(line[x]) == "." {
				continue
			}
			
			if !unicode.IsDigit(rune(line[x])) {
			var loc SymbolLocation
			loc.value = string(line[x])
			loc.coords.y = y
			loc.coords.x = x
			slocs = append(slocs, loc)
			}
		}
	}

	// walk slocs, identify which parts are adjacent
	for i := 0; i < len(slocs); i++ {
		for j := 0; j < len(pnums); j++ {
			// First check Y coordinate to make sure it's in range
			vert := slocs[i].coords.y-pnums[j].coords.y 
			if vert >= -1 && vert <= 1 {
				leftside := slocs[i].coords.x - pnums[j].coords.x
				rightside := slocs[i].coords.x - (pnums[j].coords.x + pnums[j].len)
				// Then check left vs right, because this is in an array with numbers of arbirary
				// length we check left and right adjacency separately
				if leftside >= -1 && leftside <= 1 {
					// fmt.Println("left,\t", pnums[j], ",\t", slocs[i], ",", leftside)
					pnums[j].adjacent = true
					sum += pnums[j].value
				} else if rightside >= -1 && rightside <= 1 {
					// fmt.Println("right,\t", pnums[j], ",\t", slocs[i], ",", rightside)
					pnums[j].adjacent = true
					sum += pnums[j].value
				}
			}
		}
	}

	fmt.Println(sum)
}

func partTwo(lines []string) {
	var slocs []SymbolLocation
	var pnums []PartNumber
	sum := 0

	// rows := len(lines)
	cols := len(lines[0])

	// Scan lines for symbols, ignore periods
	for y, line := range lines {
		num := 0
		count := 0
		for x := 0; x < cols; x++ {
			n, err := strconv.Atoi(string(line[x]))
			if err == nil {
				num = (num * 10) + n
				count++

				if x >= cols -1 || !unicode.IsDigit(rune(line[x+1])){
					var pnum PartNumber
					pnum.value = num
					pnum.len = count - 1 
					pnum.coords.x = x - count + 1
					pnum.coords.y = y
					pnums = append(pnums, pnum)
					num = 0
					count = 0
				}
			}
			
			// if !unicode.IsDigit(rune(line[x])) {
			if string(line[x]) == "*" {
			var loc SymbolLocation
			loc.value = string(line[x])
			loc.coords.y = y
			loc.coords.x = x
			slocs = append(slocs, loc)
			}

			if string(line[x]) == "." {
				continue
			}
		}
	}

	// walk slocs, identify which parts are adjacent - this time we need to confirm that there are
	// two part numbers for each sloc
	for i := 0; i < len(slocs); i++ {
		var adjnums []PartNumber
		for j := 0; j < len(pnums); j++ {
			// First check Y coordinate to make sure it's in range, since part numbers don't 
			// cross rows we can check both directions at the same time
			vert := slocs[i].coords.y-pnums[j].coords.y 
			if vert >= -1 && vert <= 1 {
				leftside := slocs[i].coords.x - pnums[j].coords.x
				rightside := slocs[i].coords.x - (pnums[j].coords.x + pnums[j].len)
				// Then check left vs right, because this is in an array with numbers of arbirary
				// length we check left and right adjacency separately
				if leftside >= -1 && leftside <= 1 {
					// fmt.Println("left,\t", pnums[j], ",\t", slocs[i], ",", leftside)
					pnums[j].adjacent = true
					adjnums = append(adjnums, pnums[j])
				} else if rightside >= -1 && rightside <= 1 {
					// fmt.Println("right,\t", pnums[j], ",\t", slocs[i], ",", rightside)
					pnums[j].adjacent = true
					adjnums = append(adjnums, pnums[j])
				}
			}
		}
		if len(adjnums) > 1 {
			sum += adjnums[0].value * adjnums[1].value
		}
	}

	fmt.Println(sum)
}


func main() {
	// lines := rf.ReadFile("test1.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
	partTwo(lines)
}
