package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/csvillegas/advent-of-code-2020/util"
)

var (
	questions = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

func main() {
	fmt.Println("Day 6 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	total := 0
	groups := splitTheInput()
	for _, group := range groups {
		total += countGroup1(group)
	}
	return strconv.Itoa(total)
}

func solve2() string {
	total := 0
	groups := splitTheInput()
	for _, group := range groups {
		total += countGroup2(group)
	}
	return strconv.Itoa(total)
}

func countGroup1(group string) int {
	count := 0
	for _, q := range questions {
		if strings.Contains(group, q) {
			count++
		}
	}
	return count
}

func countGroup2(group string) int {
	count := 0
	people := strings.Split(group, "\n")
	for _, q := range questions {
		numAnsweredYes := 0
		if strings.Contains(group, q) {
			for _, person := range people {
				if strings.Contains(person, q) {
					numAnsweredYes++
				}
			}
			if numAnsweredYes == len(people)-1 {
				count++
			}
		}
	}
	return count
}

func splitTheInput() []string {
	var groups []string
	bytes := []byte(util.Input6)
	start := 0
	for i, x := range bytes {
		if x == 10 && bytes[i+1] == 10 { //&& bytes[i+2] == 10 && bytes[i+3] == 9 {
			slice := bytes[start : i+1]
			groups = append(groups, string(slice))
			start = i + 2
		}
	}
	groups = append(groups, string(bytes[start:]))
	return groups
}
