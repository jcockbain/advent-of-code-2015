package main

import (
	input "aoc2015/inpututils"
	"fmt"
	"strconv"

	"regexp"
)

var (
	numbRe = regexp.MustCompile(`-?\d+`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))
}

func Part1(filename string) int {
	inp := input.ReadRaw(filename)
	total := 0
	nums := numbRe.FindAllString(inp, -1)

	for _, n := range nums {
		if n[0] == '-' {
			total -= toInt(n[1:])
		} else {
			total += toInt(n)
		}
	}

	return total
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
