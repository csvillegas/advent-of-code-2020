package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/csvillegas/advent-of-code-2020/util"
)

type instruction struct {
	com     string
	number  int
	visited bool
}

var (
	input        = strings.Split(util.Input11, "\n")
	height       = len(input)
	width        = len(input[0])
	currentState = input
	emptySeats   = 0
	test         = "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"
)

func main() {
	fmt.Println("Day 6 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	rounds := 0
	for {
		newState := ""
		for i, row := range currentState {
			//build new state
			for j, space := range row {
				newState += determineNextSpaceState1(i, j, string(space))
			}
			if i < height-1 {
				newState += "\n"
			}
		}
		if strings.Join(currentState, "\n") == newState {
			return strconv.Itoa(strings.Count(newState, "#"))
		}
		currentState = strings.Split(newState, "\n")
		rounds++
	}
}

func solve2() string {
	for {
		newState := ""
		emptySeats = 0
		for i, row := range currentState {
			//build new state
			for j, space := range row {
				newState += determineNextSpaceState2(i, j, string(space))
			}
			if i < height-1 {
				newState += "\n"
			}
		}
		if strings.Join(currentState, "\n") == newState {
			return strconv.Itoa(strings.Count(newState, "#"))
		}
		currentState = strings.Split(newState, "\n")
	}
}

func determineNextSpaceState1(i, j int, space string) string {
	adjacentSeats := findAdjacentSeats1(i, j)
	if space == "L" {
		if !strings.Contains(adjacentSeats, "#") {
			return "#"
		} else {
			emptySeats++
			return space // no change
		}
	}
	if space == "#" {
		if strings.Count(adjacentSeats, "#") >= 4 {
			emptySeats++
			return "L"
		} else {
			return space // no change
		}
	}
	return "."
}

func determineNextSpaceState2(i, j int, space string) string {
	if space == "." {
		return space
	}
	adjacentSeats := findAdjacentSeats2(i, j)
	if space == "L" {
		if !strings.Contains(adjacentSeats, "#") {
			return "#"
		} else {
			return space // no change
		}
	}
	if space == "#" {
		if strings.Count(adjacentSeats, "#") >= 5 {
			return "L"
		} else {
			return space // no change
		}
	}
	return "."
}

func findAdjacentSeats1(i, j int) string {
	seats := ""
	lj := j - 1
	if lj >= 0 { // left
		seats += string(currentState[i][lj])
	}
	rj := j + 1
	if rj < width { // right
		seats += string(currentState[i][rj])
	}
	ui := i - 1
	if ui >= 0 { // up
		seats += string(currentState[ui][j])
	}
	di := i + 1
	if di < height { // down
		seats += string(currentState[di][j])
	}
	if lj >= 0 && ui >= 0 { // NE
		seats += string(currentState[ui][lj])
	}
	if lj >= 0 && di < height { // SE
		seats += string(currentState[di][lj])
	}
	if rj < width && ui >= 0 { // NW
		seats += string(currentState[ui][rj])
	}
	if rj < width && di < height { // SW
		seats += string(currentState[di][rj])
	}
	return seats
}

func findAdjacentSeats2(i, j int) string {
	seats := ""
	lj := j - 1
	if lj >= 0 { // left
		left := string(currentState[i][lj])
		for left == "." && lj > 0 {
			lj--
			left = string(currentState[i][lj])
		}
		seats += string(currentState[i][lj])
	}
	rj := j + 1
	if rj < width { // right
		right := string(currentState[i][rj])
		for right == "." && rj < width-1 {
			rj++
			right = string(currentState[i][rj])
		}
		seats += string(currentState[i][rj])
	}
	ui := i - 1
	if ui >= 0 { // up
		up := string(currentState[ui][j])
		for up == "." && ui > 0 {
			ui--
			up = string(currentState[ui][j])
		}
		seats += string(currentState[ui][j])
	}
	di := i + 1
	if di < height { // down
		down := string(currentState[di][j])
		for down == "." && di < height-1 {
			di++
			down = string(currentState[di][j])
		}
		seats += string(currentState[di][j])
	}
	lj = j - 1
	ui = i - 1
	if lj >= 0 && ui >= 0 { // NW
		nw := string(currentState[ui][lj])
		for nw == "." && lj > 0 && ui > 0 {
			lj--
			ui--
			nw = string(currentState[ui][lj])
		}
		seats += string(currentState[ui][lj])
	}
	lj = j - 1
	di = i + 1
	if lj >= 0 && di < height { // SW
		sw := string(currentState[di][lj])
		for sw == "." && lj > 0 && di < height-1 {
			lj--
			di++
			sw = string(currentState[di][lj])
		}
		seats += string(currentState[di][lj])
	}
	rj = j + 1
	ui = i - 1
	if rj < width && ui >= 0 { // NE
		ne := string(currentState[ui][rj])
		for ne == "." && rj < width-1 && ui > 0 {
			rj++
			ui--
			ne = string(currentState[ui][rj])
		}
		seats += string(currentState[ui][rj])
	}
	rj = j + 1
	di = i + 1
	if rj < width && di < height { // SE
		se := string(currentState[di][rj])
		for se == "." && rj < width-1 && di < height-1 {
			rj++
			di++
			se = string(currentState[di][rj])
		}
		seats += string(currentState[di][rj])
	}
	return seats
}

func sanitizeKey(k string) string {
	return strings.Join(strings.Fields(k), "")
}
