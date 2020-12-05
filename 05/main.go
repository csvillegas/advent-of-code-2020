package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/csvillegas/advent-of-code-2020/util"
)

var (
	boardingCodes = strings.Split(util.Input5, "\n")
)

func main() {
	fmt.Println("Day 5 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	largest := 0
	for _, x := range boardingCodes {
		r := findRow(x[:7])
		c := findCol(x[7:])

		code := r*8 + c
		if code > largest {
			largest = code
		}
	}
	return strconv.Itoa(largest)
}

func solve2() string {
	var codes []int
	for _, x := range boardingCodes {
		r := findRow(x[:7])
		c := findCol(x[7:])
		code := r*8 + c
		codes = append(codes, code)
	}
	sort.Ints(codes)
	for i, _ := range codes {
		id := codes[i]
		if i+1 != len(codes) && id+1 != codes[i+1] {
			return strconv.Itoa(id + 1)
		}
	}
	return ""
}

func findRow(s string) int {
	f := 0
	b := 127
	for _, x := range s {
		char := string(x)
		if char == "F" {
			b = f + (b-f)/2
		}
		if char == "B" {
			f = b - (b-f)/2
		}
	}
	// if f != b {
	// 	return err
	// }
	return f
}

func findCol(s string) int {
	l := 0
	r := 7
	for _, x := range s {
		char := string(x)
		if char == "L" {
			r = l + (r-l)/2
		}
		if char == "R" {
			l = r - (r-l)/2
		}
	}
	// if r != l {
	// 	return err
	// }
	return l
}
