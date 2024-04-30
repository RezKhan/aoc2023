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
		for x := 0; x < cols; x++ {
			// parse numbers and add to a slice with locations
			if unicode.IsDigit(rune(line[x])) {
				num := 0
				var t []int
				var pnum PartNumber
				tmpx := x 
				for i := x; i < cols; i++ {
					n, err := strconv.Atoi(string(line[i]))
					if err == nil {
						t = append(t, n)
						num = (num * 10) + n
					}  else {
						break
					}
					x = i
				}
				pnum.value = num
				pnum.len = len(t) - 1
				pnum.coords.x = tmpx
				pnum.coords.y = y
				pnums = append(pnums, pnum)
				if (x + 1) < cols {
					x += 1
				}
			}
					
					
			// ignore periods "."
			if string(line[x]) == "." {
				continue
			}
			// n, err := strconv.Atoi(string(line[x]))
			// if err == nil {
			// 	t := n
			// 	for i := 1; x+i < cols; i++ {
			// 		fmt.Println(t, y, x, ":", x+i)
			// 		n, err = strconv.Atoi(string(line[x+i]))
			// 		if err == nil {
			// 			t = (t * 10) + n
			// 		} else {
			// 			var pnum PartNumber
			// 			pnum.value = t
			// 			pnum.coords.y = y
			// 			pnum.coords.x = x
			// 			pnum.len = i - 1
			// 			pnums = append(pnums, pnum)
			// 			x += i - 1
			// 			break
			// 		}
			// 	}
			// 	continue
			// }

			var loc SymbolLocation
			loc.value = string(line[x])
			loc.coords.y = y
			loc.coords.x = x
			slocs = append(slocs, loc)
		}
	}

	for _, sloc := range slocs {
		_, err := strconv.Atoi(sloc.value)
		if err == nil {
			fmt.Println(sloc)
		}
	}

	// walk slocs, identify which parts are adjacent
	for i := 0; i < len(slocs); i++ {
		for j := 0; j < len(pnums); j++ {
			// y position validation: WORKING
			if slocs[i].coords.y-pnums[j].coords.y >= -1 && slocs[i].coords.y-pnums[j].coords.y <= 1 {
				// fmt.Println("values: ", slocs[i], pnums[j])
				// fmt.Println(pnums[j], "left side:", slocs[i].coords.x - pnums[j].coords.x, "right side:", slocs[i].coords.x - (pnums[j].coords.x + pnums[j].len))
				leftside := slocs[i].coords.x - pnums[j].coords.x
				rightside := slocs[i].coords.x - (pnums[j].coords.x + pnums[j].len)
				// fmt.Println(pnums[j], leftside, rightside)
				if leftside >= -1 && leftside <= 1 {
					pnums[j].adjacent = true
					fmt.Println("left,\t", pnums[j], ",\t", slocs[i], ",", leftside)
				} else if rightside >= -1 && rightside <= 1 {
					pnums[j].adjacent = true
					fmt.Println("right,\t", pnums[j], ",\t", slocs[i], ",", rightside)
				}
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	// lines := rf.ReadFile("test1.txt")
	lines := rf.ReadFile("input.txt")
	partOne(lines)
}
