package main

import (
	input "aoc2015/inpututils"

	"fmt"
	"strings"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

type loc struct {
	x, y int
}

func Part1(filename string) int {
	pos := loc{0, 0}
	dirs := input.ReadRaw(filename)

	delivered := map[loc]int{}

	for _, d := range strings.Split(dirs, "") {
		delivered[pos] += 1
		pos = move(pos, d)
	}

	return len(delivered)
}

func Part2(filename string) int {
	pos := loc{0, 0}
	roboPos := loc{0, 0}
	dirs := input.ReadRaw(filename)

	delivered := map[loc]int{pos: 2}

	for i, d := range strings.Split(dirs, "") {
		if i%2 == 0 {
			pos = move(pos, d)
			delivered[pos] += 1
		} else {
			roboPos = move(roboPos, d)
			delivered[roboPos] += 1
		}
	}

	return len(delivered)
}

func move(pos loc, dir string) loc {
	if dir == string("^") {
		pos.y = pos.y + 1
	}
	if dir == string("v") {
		pos.y = pos.y - 1
	}
	if dir == string(">") {
		pos.x = pos.x + 1
	}
	if dir == string("<") {
		pos.x = pos.x - 1
	}
	return pos
}
