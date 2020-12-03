package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/csvillegas/advent-of-code-2020/util"
)

var policies = strings.Split(util.Input2, "\n")

func main() {
	fmt.Println("Day 2 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	const (
		ranges = 0
		letter = 1
		pw     = 2
	)
	var numValid = 0
	for _, str := range policies {
		policy := strings.Split(str, " ")
		count := strings.Count(policy[pw], string((policy[letter])[0]))
		rangeValues := strings.Split(policy[ranges], "-")
		floor, _ := strconv.Atoi(rangeValues[0])
		ciel, _ := strconv.Atoi(rangeValues[1])
		if count >= floor && count <= ciel {
			numValid++
		}
	}
	return strconv.Itoa(numValid)
}

func solve2() string {
	const (
		ranges = 0
		letter = 1
		pw     = 2
	)
	var numValid = 0
	for _, str := range policies {
		policy := strings.Split(str, " ")
		char := (policy[letter])[0]
		p := policy[pw]
		rangeValues := strings.Split(policy[ranges], "-")
		pos1, _ := strconv.Atoi(rangeValues[0])
		pos2, _ := strconv.Atoi(rangeValues[1])
		pos1--
		pos2--
		if pos1 != pos2 && ((p[pos1] == char && p[pos2] != char) || (p[pos2] == char && p[pos1] != char)) {
			numValid++
		}
	}
	return strconv.Itoa(numValid)
}
