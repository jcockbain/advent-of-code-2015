package main

import (
	input "aoc2015/inpututils"

	"fmt"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

func Part1(filename string) int {
	lines := input.ReadRaw(filename)
	s := 0
	for _, c := range lines {
		if string(c) == "(" {
			s += 1
		} else if string(c) == ")" {
			s -= 1
		}
	}

	return s
}

func Part2(filename string) int {
	lines := input.ReadRaw(filename)

	s := 0

	for i, c := range lines {
		if string(c) == "(" {
			s += 1
		} else if string(c) == ")" {
			if s == 0 {
				return i + 1
			}
			s -= 1
		}
	}

	return s
}
