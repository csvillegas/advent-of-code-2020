package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/csvillegas/advent-of-code-2020/util"
)

var (
	skiSlopes = strings.Split(util.Input3, "\n")
	columns   = len(skiSlopes[0])
	tcCh      = make(chan int)
	tcCh1     = make(chan int)
	tcCh2     = make(chan int)
	tcCh3     = make(chan int)
	tcCh4     = make(chan int)
	tcCh5     = make(chan int)
	chanels   = []chan int{tcCh, tcCh1, tcCh2, tcCh3, tcCh4, tcCh5}
	slopes    = []int{3, 1, 3, 5, 7, 1}
)

func main() {
	fmt.Println("Day 3 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	go boring(0, 1)
	return strconv.Itoa(<-chanels[0])
}

func solve2() string {
	go boring(1, 1)
	go boring(2, 1)
	go boring(3, 1)
	go boring(4, 1)
	go boring(5, 2)
	return strconv.Itoa(<-chanels[1] * <-chanels[2] * <-chanels[3] * <-chanels[4] * <-chanels[5])
}

func boring(i int, step int) {
	var treeCount = 0
	var x = 0
	for row, _ := range skiSlopes {
		if row%step == 0 {
			treeCount += checkNextSpace(x, row)
			x += slopes[i]
			if x >= columns {
				x = x - columns
			}
		}
	}
	chanels[i] <- treeCount
}

func checkNextSpace(x int, y int) int {
	treeCode := ([]byte("#"))[0]
	if skiSlopes[y][x] == treeCode {
		return 1
	} else {
		return 0
	}
}
