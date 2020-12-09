package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/csvillegas/advent-of-code-2020/util"
)

const pre = 25

var (
	input     = strings.Split(util.Input9, "\n")
	solution1 = make(chan int)
	target    = make(chan int)
	position  = make(chan int)
	solution2 = make(chan int)
)

func main() {
	fmt.Println("Day 6 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	for i, x := range input {
		if i >= pre {
			go checkIfValid(i, x, input[i-pre:i])
		}
	}
	return strconv.Itoa(<-solution1)
}

func checkIfValid(pos int, n string, a []string) {
	val, _ := strconv.Atoi(n)
	for i, x := range a {
		v1, _ := strconv.Atoi(x)
		for _, y := range a[i+1:] {
			v2, _ := strconv.Atoi(y)
			if v1+v2 == val && v1 != v2 {
				return
			}
		}
	}
	solution1 <- val
	target <- val
	position <- pos
	return
}

func solve2() string {
	target := <-target
	pos := <-position
	for i, _ := range input[:pos] {
		go findWeakness(target, input[i:pos])
	}
	return strconv.Itoa(<-solution2)
}

func findWeakness(target int, a []string) {
	largest := 0
	first, _ := strconv.Atoi(a[0])
	smallest := first
	sum := 0
	for _, x := range a {
		xVal, _ := strconv.Atoi(x)
		if xVal > largest {
			largest = xVal
		}
		if xVal < smallest {
			smallest = xVal
		}
		sum += xVal
		if sum == target {
			solution2 <- largest + smallest
		}
	}
	return
}
