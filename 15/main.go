package main

import (
	"fmt"
	"strconv"
)

var (
	input = []int{9, 3, 1, 0, 8, 4}
)

func main() {
	fmt.Println("Day 3 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	return solve(2020)
}

func solve2() string {
	return solve(30000000)
}

func solve(nth int) string {
	spoken := input
	spokenBefore := map[int]int{}
	for i, n := range spoken {
		if i < len(spoken)-1 {
			spokenBefore[n] = i + 1
		}
	}
	turn := len(spoken) + 1
	lastSpoken := spoken[turn-2]
	for turn <= nth {
		lastTurn := turn - 1
		if val, ok := spokenBefore[lastSpoken]; ok {
			if lastTurn == val {
				spoken = append(spoken, 1)
			} else {
				spoken = append(spoken, lastTurn-val)
			}
		} else {
			//never spoken
			spoken = append(spoken, 0)
		}
		spokenBefore[lastSpoken] = turn - 1
		lastSpoken = spoken[turn-1]
		turn++
	}
	return strconv.Itoa(lastSpoken)
}
