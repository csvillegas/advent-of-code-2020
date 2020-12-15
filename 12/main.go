package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/csvillegas/advent-of-code-2020/util"
)

var (
	input      = strings.Split(util.Input12, "\n")
	tcCh       = make(chan int)
	tcCh1      = make(chan int)
	tcCh2      = make(chan int)
	tcCh3      = make(chan int)
	tcCh4      = make(chan int)
	tcCh5      = make(chan int)
	chanels    = []chan int{tcCh, tcCh1, tcCh2, tcCh3, tcCh4, tcCh5}
	test       = "F10\nN3\nF7\nR90\nF11"
	directions = []string{"E", "S", "W", "N"} // clockwise order (to right)
)

func main() {
	fmt.Println("Day 3 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	x := 0  // east-west
	y := 0  // north-south
	dh := 0 // direction-head

	for _, c := range input {
		com := string(c[:1])
		val, _ := strconv.Atoi(string(c[1:]))
		switch com {
		case "N":
			y += val
		case "S":
			y -= val
		case "E":
			x += val
		case "W":
			x -= val
		case "R":
			rot := val / 90
			if rot >= 4 {
				dh = (dh + (rot % 4)) % 4
			} else {
				dh = (dh + rot) % 4
			}
		case "L":
			rot := val / 90
			if rot >= 4 {
				dh = dh - (rot % 4)
			} else {
				dh -= rot
			}
			if dh == -1 {
				dh = 3
			}
			if dh == -2 {
				dh = 2
			}
			if dh == -3 {
				dh = 1
			}
		case "F":
			direction := directions[dh]
			switch direction {
			case "N":
				y += val
			case "S":
				y -= val
			case "E":
				x += val
			case "W":
				x -= val
			}
		}
	}
	return strconv.Itoa(int(math.Abs(float64(x)) + math.Abs(float64(y))))
}

func solve2() string {
	wx := 10
	wy := 1
	x := 0 // east-west
	y := 0 // north-south

	for _, c := range input {
		com := string(c[:1])
		val, _ := strconv.Atoi(string(c[1:]))
		switch com {
		case "N":
			wy += val
		case "S":
			wy -= val
		case "E":
			wx += val
		case "W":
			wx -= val
		case "R":
			var moves int
			rot := val / 90
			if rot >= 4 {
				moves = rot % 4
			} else {
				moves = rot
			}
			wx, wy = rotateR(moves, wx, wy)
		case "L":
			var moves int
			rot := val / 90
			if rot >= 4 {
				moves = rot % 4
			} else {
				moves = rot
			}
			wx, wy = rotateL(moves, wx, wy)
		case "F":
			x += wx * val
			y += wy * val
		}
	}
	return strconv.Itoa(int(math.Abs(float64(x)) + math.Abs(float64(y))))
}

func rotateR(moves, wx, wy int) (int, int) {
	for i := 1; i <= moves; i++ {

		tempX := wx
		tempY := wy

		if wx > 0 && wy > 0 { // pp
			wx = tempY  // pos
			wy = -tempX // neg
		} else if wx > 0 && wy < 0 { // pn
			wx = tempY  // neg
			wy = -tempX // neg
		} else if wx < 0 && wy > 0 { // np
			wx = tempY  // pos
			wy = -tempX // pos
		} else if wx < 0 && wy < 0 { // nn
			wx = tempY  // neg
			wy = -tempX //pos
		}
	}
	return wx, wy
}

func rotateL(moves, wx, wy int) (int, int) {
	for i := 1; i <= moves; i++ {

		tempX := wx
		tempY := wy

		if wx > 0 && wy > 0 { // pp
			wx = -tempY // neg
			wy = tempX  // pos
		} else if wx > 0 && wy < 0 { // pn
			wx = -tempY // pos
			wy = tempX  // pos
		} else if wx < 0 && wy > 0 { // np
			wx = -tempY // neg
			wy = tempX  // neg
		} else if wx < 0 && wy < 0 { // nn
			wx = -tempY // pos
			wy = tempX  //neg
		}
	}
	return wx, wy
}
