package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/csvillegas/advent-of-code-2020/util"
)

var ledger []string

var solutionCh1 = make(chan int)
var solutionCh2 = make(chan int)

func main() {
	// input := inputs.Input1
	ledger = strings.Split(util.Input1, "\n")
	fmt.Println("Day 1 Advent Of Code")
	solve1()
	fmt.Println("Puzzle 1: " + strconv.Itoa(<-solutionCh1))
	solve2()
	fmt.Println("Puzzle 2: " + strconv.Itoa(<-solutionCh2))
}

func solve1() {
	for i, x := range ledger {
		go boring1(i, -1, x)
	}
	return
}

func solve2() {
	for i, x := range ledger {
		go boring2(i, x)
	}
}

func boring1(i int, j int, x string) {
	candidate, _ := strconv.Atoi(x)
	for k, y := range ledger {
		match, _ := strconv.Atoi(y)
		if candidate+match == 2020 && i != k && i != j {
			if j == -1 {
				solutionCh1 <- candidate * match
			} else {
				v1, _ := strconv.Atoi(ledger[i])
				v2, _ := strconv.Atoi(ledger[j])
				v3, _ := strconv.Atoi(ledger[k])
				solutionCh2 <- v1 * v2 * v3
			}
		}
	}
}

func boring2(i int, x string) {
	candidate, _ := strconv.Atoi(x)
	for j, y := range ledger {
		match, _ := strconv.Atoi(y)
		value := strconv.Itoa(candidate + match)
		go boring1(i, j, value)
	}
}
