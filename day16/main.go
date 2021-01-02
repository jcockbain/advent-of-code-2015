package main

import (
	input "aoc2015/inpututils"
	"fmt"

	"regexp"
	"strconv"
	"strings"
)

var (
	re = regexp.MustCompile(`[a-z]*: (\d+)`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

func Part1(filename string) []int {
	lines := input.ReadLines(filename)

	sue := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
	}
	possibles := []int{}
	for i, line := range lines {
		validSue := true
		sueOwns := re.FindAllString(line, -1)
		for _, poss := range sueOwns {
			split := strings.Split(poss, ": ")
			thing, number := split[0], toInt(split[1])
			if sue[thing] != number {
				validSue = false
				break
			}
		}
		if validSue {
			possibles = append(possibles, i+1)
		}
	}
	return possibles
}

func Part2(filename string) []int {
	lines := input.ReadLines(filename)

	sue := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	possibles := []int{}
	for i, line := range lines {
		validSue := true
		sueOwns := re.FindAllString(line, -1)
		for _, poss := range sueOwns {
			split := strings.Split(poss, ": ")
			thing, number := split[0], toInt(split[1])
			if thing == "cats" || thing == "trees" {
				if number <= sue[thing] {
					validSue = false
					break
				}
			} else if thing == "pomeranians" || thing == "goldfish" {
				if number >= sue[thing] {
					validSue = false
					break
				}
			} else {
				if sue[thing] != number {
					validSue = false
					break
				}
			}
		}
		if validSue {
			possibles = append(possibles, i+1)
		}
	}
	return possibles
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
