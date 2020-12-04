package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/csvillegas/advent-of-code-2020/util"
)

var (
	tcCh    = make(chan int)
	tcCh1   = make(chan int)
	tcCh2   = make(chan int)
	tcCh3   = make(chan int)
	tcCh4   = make(chan int)
	tcCh5   = make(chan int)
	chanels = []chan int{tcCh, tcCh1, tcCh2, tcCh3, tcCh4, tcCh5}

	test = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
	hcl:#623a2f
	
	eyr:2029 ecl:blu cid:129 byr:1989
	iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm
	
	hcl:#888785
	hgt:164cm byr:2001 iyr:2015 cid:88
	pid:545766238 ecl:hzl
	eyr:2022
	
	iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

	passportFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	optionalField  = passportFields[7]
)

func main() {
	fmt.Println("Day 4 Advent Of Code")

	fmt.Println("Puzzle 1: " + solve1())
	fmt.Println("Puzzle 2: " + solve2())
}

func solve1() string {
	passports := splitTheInput()
	var valid = 0
	for _, pp := range passports {
		if strings.Contains(pp, passportFields[0]) &&
			strings.Contains(pp, passportFields[1]) &&
			strings.Contains(pp, passportFields[2]) &&
			strings.Contains(pp, passportFields[3]) &&
			strings.Contains(pp, passportFields[4]) &&
			strings.Contains(pp, passportFields[5]) &&
			strings.Contains(pp, passportFields[6]) {
			valid++
		}
	}
	return strconv.Itoa(valid)
}

func solve2() string {
	passports := splitTheInput()
	validPassorts := 0
	for _, pp := range passports {
		validFields := 0
		fields := strings.Split(strings.Replace(pp, "\n", " ", -1), " ")
		for _, f := range fields {
			validFields += validateField(f)
		}
		if validFields >= 7 &&
			strings.Contains(pp, passportFields[0]) &&
			strings.Contains(pp, passportFields[1]) &&
			strings.Contains(pp, passportFields[2]) &&
			strings.Contains(pp, passportFields[3]) &&
			strings.Contains(pp, passportFields[4]) &&
			strings.Contains(pp, passportFields[5]) &&
			strings.Contains(pp, passportFields[6]) {
			// fmt.Println("valid", pp)
			validPassorts++
		}
	}
	return strconv.Itoa(validPassorts)
}

func splitTheInput() []string {
	var pps []string
	bytes := []byte(util.Input4)
	// bytes := []byte(test)

	//fmt.Println(bytes)
	start := 0
	for i, x := range bytes {
		if x == 10 && bytes[i+1] == 10 { //&& bytes[i+2] == 10 && bytes[i+3] == 9 {
			slice := bytes[start : i+1]
			pps = append(pps, string(slice))
			start = i + 2
		}
	}
	pps = append(pps, string(bytes[start:]))
	return pps
}

func validateField(f string) int {
	if strings.Contains(f, ":") {
		pair := strings.Split(f, ":")
		pair[0] = sanitizeKey(pair[0])
		pair[1] = sanitizeKey(pair[1])

		// fmt.Println(pair[0])
		// fmt.Println(pair[1])

		switch pair[0] {
		case "byr":
			value, _ := strconv.Atoi(pair[1])
			if len(pair[1]) == 4 &&
				value >= 1920 &&
				value <= 2002 {
				// fmt.Println("valid byr")
				return 1
			}
			// fmt.Println("invalid byr")
		case "iyr":
			value, _ := strconv.Atoi(pair[1])
			if len(pair[1]) == 4 &&
				value >= 2010 &&
				value <= 2020 {
				// fmt.Println("valid iyr")
				return 1
			}
			// fmt.Println("invalid iyr")

		case "eyr":
			value, _ := strconv.Atoi(pair[1])
			if len(pair[1]) == 4 &&
				value >= 2020 &&
				value <= 2030 {
				// fmt.Println("valid eyr")

				return 1
			}
			// fmt.Println("invalid eyr")

		case "hgt":
			height := pair[1]
			num := height[:len(height)-2]
			value, err := strconv.Atoi(num)
			if err != nil {
				// fmt.Println("invalid hgt 1")
				return 0
			}

			if strings.Contains(height, "cm") {
				if value >= 150 && value <= 193 {
					// fmt.Println("valid hgt cm")
					return 1
				}
			} else if strings.Contains(height, "in") {
				if value >= 59 && value <= 76 {
					// fmt.Println("valid hgt in")
					return 1
				}
			}
			// fmt.Println("invalid hgt 2", height, num)

			return 0
		case "hcl":
			_, err := strconv.ParseUint(strings.Replace(pair[1], "#", "0x", -1), 0, 64)
			if err != nil {
				// pair[1] is  not a valid hex value
				// fmt.Println("invalid hcl")

				return 0
			}
			// fmt.Println("valid hcl")

			return 1
		case "ecl":
			ec := pair[1]
			if ec == "amb" ||
				ec == "blu" ||
				ec == "brn" ||
				ec == "gry" ||
				ec == "grn" ||
				ec == "hzl" ||
				ec == "oth" {
				// fmt.Println("valid ecl")

				return 1
			}
			// fmt.Println("invalid ecl")

		case "pid":
			if len(pair[1]) == 9 {
				// fmt.Println("valid pid")

				return 1
			}
			// fmt.Println("invalid pid")

		case "cid":
			// do nothing
		}
	}
	return 0
}

func sanitizeKey(k string) string {
	return strings.Join(strings.Fields(k), "")
}
