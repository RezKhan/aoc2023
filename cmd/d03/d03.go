package main

import (
	"fmt"
	rf "readfile"
	"strconv"
	"strings"
)
type Coords struct {
	y int
	x int
}
type PartNumber struct {
	value int
	coords Coords
	len int
	adjacent bool
}

type SymbolLocation struct {
	value string
	coords Coords
}


func partOne(lines []string) {
	var s string
	var slocs []SymbolLocation
	var pnums []PartNumber
	sum := 0

	// rows := len(lines)
	cols := len(lines[0])
	
	// Scan lines for symbols, ignore periods
	for y, line := range lines {
		for x := 0; x < cols; x++ {
			// ignore periods "."
			if string(line[x]) == "." {
				continue
			}
			// parse numbers and add to a slice with locations
			n, err := strconv.Atoi(string(line[x])) 
			if err == nil {
				t := n
				for i := 1; i < cols; i++ {
					if x + i < cols {
						n, err = strconv.Atoi(string(line[x+i]))
						if err == nil {
							t = (t * 10) + n
						} else {
							var pnum PartNumber
							pnum.value = t
							pnum.coords.y = y
							pnum.coords.x = x
							pnum.len = i - 1
							pnums = append(pnums, pnum)
							x += i - 1
							break
						}
					}
				}
				continue
			}

			var loc SymbolLocation
			loc.value = string(line[x])
			loc.coords.y = y
			loc.coords.x = x
			slocs = append(slocs, loc)
			
			if !strings.Contains(s, string(line[x])) {
				s += string(line[x])
			}
		}
	}
	// walk slocs, identify which parts are adjacent
	for i := 0; i < len(slocs); i++ {
		for j := 0; j < len(pnums); j++ {
			// y position validation: WORKING
			if slocs[i].coords.y - pnums[j].coords.y >= -1 && slocs[i].coords.y - pnums[j].coords.y <= 1 {
				// fmt.Println("values: ", slocs[i], pnums[j])
				// fmt.Println(pnums[j], "right side:", slocs[i].coords.x - (pnums[j].coords.x + pnums[j].len))
				// fmt.Println("right side:", slocs[i].coords.x - pnums[j].coords.x)
				if slocs[i].coords.x - (pnums[j].coords.x + pnums[j].len) <= 1 && slocs[i].coords.x - (pnums[j].coords.x + pnums[j].len) >= -1  {
					pnums[j].adjacent = true
					sum += pnums[j].value
				} else if slocs[i].coords.x - pnums[j].coords.x >= -1 && slocs[i].coords.x - pnums[j].coords.x <= 1 {
					pnums[j].adjacent = true
					sum += pnums[j].value
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
