package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/csvillegas/advent-of-code-2020/util"
)

var (
	input   = strings.Split(util.Input14, "\n")
	memory  = map[string]int64{}
	memory2 = map[int64]int64{}
)

func main() {
	fmt.Println("Day 3 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	var mask string
	for _, x := range input {
		pair := strings.Split(x, " = ")
		if pair[0] == "mask" {
			mask = pair[1]
		} else {
			mem := pair[0] //address
			val, _ := strconv.Atoi(pair[1])
			valMasked, _ := applyBitmask(mask, int64(val))
			memory[mem] = valMasked
		}
	}
	sum := int64(0)
	for _, val := range memory {
		sum += val
	}
	return strconv.Itoa(int(sum))
}

func applyBitmask(mask string, number int64) (int64, error) {
	var masked string
	nBi := strconv.FormatInt(number, 2)
	nBi = correctSize(nBi, len(mask))
	masked = nBi
	for i, x := range mask {
		if string(x) != "X" {
			masked = replaceAtIndex(masked, x, i)
		}
	}
	return strconv.ParseInt(string(masked), 2, 64)
}

func correctSize(bi string, s int) string {
	for len(bi) < s {
		bi = "0" + bi
	}
	return bi
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func solve2() string {
	var mask string
	for _, x := range input {
		pair := strings.Split(x, " = ")
		if pair[0] == "mask" {
			mask = pair[1]
		} else {
			mem := pair[0][4 : len(pair[0])-1] // address-string
			memVal, _ := strconv.Atoi(mem)     // address-int
			val, _ := strconv.Atoi(pair[1])    // value to write
			addresses := applyBitmaskQ(mask, int64(memVal))
			for _, a := range addresses {
				memory2[a] = int64(val)
			}
		}
	}
	sum := int64(0)
	for _, val := range memory2 {
		sum += val
	}
	return strconv.Itoa(int(sum))
}

func applyBitmaskQ(mask string, number int64) []int64 {
	var masked string
	var posArr []int
	nBi := strconv.FormatInt(number, 2) // address as bi-string
	nBi = correctSize(nBi, len(mask))   // match size w/ leading 0's
	masked = nBi
	for i, x := range mask {
		if string(x) == "1" || string(x) == "X" {
			masked = replaceAtIndex(masked, x, i)
			if string(x) == "X" {
				posArr = append(posArr, i) // note where X's are
			}
		}
	}
	biArr := []string{} // array of binary strings (0-numAdd)
	numX := strings.Count(masked, "X")
	numAdd := math.Pow(2, float64(numX)) // number of possible bi strings
	for i := 0; float64(i) < numAdd; i++ {
		biArr = append(biArr, correctSize(strconv.FormatInt(int64(i), 2), numX))
	}
	var addresses = []int64{}
	for _, x := range biArr {
		address := masked
		for i, d := range x {
			address = replaceAtIndex(address, d, posArr[i])
		}
		t, _ := strconv.ParseInt(string(address), 2, 64)
		addresses = append(addresses, t)
	}
	return addresses
}
