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

const (
	nop = "nop"
	acc = "acc"
	jmp = "jmp"
)

var (
	input = strings.Split(util.Input8, "\n")
	ch    = make(chan int)
)

func main() {
	fmt.Println("Day 6 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	accumulator := 0
	readHead := 0
	program := createIntructionList()
	for {
		instruction := program[readHead]
		if instruction.visited {
			// solution
			return strconv.Itoa(accumulator)
		}
		instruction.visited = true
		program[readHead] = instruction
		switch instruction.com {
		case acc:
			accumulator += instruction.number
			readHead++
		case jmp:
			readHead += instruction.number
		case nop:
			readHead++
		}
	}
}

func solve2() string {
	program := createIntructionList()
	for i, instruction := range program {
		if instruction.com == nop || instruction.com == jmp {
			go haltingProblem(i)
		}
	}
	return strconv.Itoa(<-ch)
}

func createIntructionList() []instruction {
	var program []instruction
	for _, w := range input {
		pair := strings.Split(w, " ")
		n, _ := strconv.Atoi(pair[1])
		instruction := instruction{
			com:     pair[0],
			number:  n,
			visited: false,
		}
		program = append(program, instruction)
	}
	return program
}

func haltingProblem(i int) {
	program := createIntructionList()
	program[i].com = swap(program[i].com)
	accumulator := 0
	readHead := 0
	for {
		instruction := program[readHead]
		if instruction.visited {
			// infinite loop
			return
		}
		instruction.visited = true
		program[readHead] = instruction
		switch instruction.com {
		case acc:
			accumulator += instruction.number
			readHead++
		case jmp:
			readHead += instruction.number
		case nop:
			readHead++
		}
		if readHead == len(program) {
			// solution
			ch <- accumulator
			return
		}
	}
}

func swap(s string) string {
	if s == nop {
		return jmp
	} else {
		return nop
	}
}
