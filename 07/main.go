package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/csvillegas/advent-of-code-2020/util"
)

type colorBagRule struct {
	color  string
	number int
}

type colorBagRules []colorBagRule

type colorBag map[string]colorBagRules

const empty = "nootherbags"
const target = "shinygoldbags"

var (
	input = strings.Split(util.Input7, "\n")
)

func main() {
	fmt.Println("Day 6 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	var bags = makeCB()
	var count = 0
	for key, _ := range bags {
		if key != target {
			count += findGoldBag(key, bags)
		}
	}
	return strconv.Itoa(count)
}

func solve2() string {
	var bags = makeCB()
	total := countBagsNeeded(bags, target, 1)
	return strconv.Itoa(total - 1)
}

func makeCB() colorBag {
	var bags = make(colorBag)
	for _, x := range input {
		sx := sanitizeKey(x)
		colorAndRules := strings.Split(sx, "contain")
		newColor := colorAndRules[0]
		bags[newColor] = colorBagRules{}
		colorAndRules[1] = strings.Replace(colorAndRules[1], ".", "", -1)
		rules := strings.Split(colorAndRules[1], ",")
		for _, r := range rules {
			if r != empty {
				if string(r[len(r)-1]) != "s" {
					// make all colors consistent as plurals
					r += "s"
				}
				n, _ := strconv.Atoi(r[:1])
				newRule := colorBagRule{
					color:  r[1:],
					number: n,
				}
				bags[newColor] = append(bags[newColor], newRule)
			}
		}
	}
	return bags
}

func findGoldBag(key string, bags colorBag) int {
	for _, rule := range bags[key] {
		if rule.color == target {
			return 1
		}
		n := findGoldBag(rule.color, bags)
		if n > 0 {
			return n
		}
	}
	// no other bags
	return 0
}

func countBagsNeeded(bags colorBag, key string, multiplier int) int {
	counts := []int{}
	for _, rule := range bags[key] {
		if rule.color != empty {
			n := countBagsNeeded(bags, rule.color, rule.number)
			counts = append(counts, n*multiplier)
		}
	}
	if len(counts) == 0 {
		return multiplier
	}
	sum := 0
	for _, c := range counts {
		sum += c
	}
	return sum + multiplier
}

func sanitizeKey(k string) string {
	return strings.Join(strings.Fields(k), "")
}
