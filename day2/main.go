package main

import (
	input "aoc2015/inpututils"
	"strconv"

	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

func Part1(filename string) int {
	lines := input.ReadLines(filename)
	total := 0
	for _, line := range lines {
		total += PaperNeeded(line)
	}
	return total
}

func Part2(filename string) int {
	lines := input.ReadLines(filename)
	total := 0
	for _, line := range lines {
		total += RibbonNeeded(line)
	}
	return total
}

func PaperNeeded(dim string) int {
	dims := strings.Split(dim, "x")
	l, w, h := toInt(dims[0]), toInt(dims[1]), toInt(dims[2])
	surfaceArea := (2 * l * w) + (2 * w * h) + (2 * h * l)
	sorted := []int{l, w, h}
	sort.Ints(sorted)
	smallestSide := sorted[0] * sorted[1]
	return surfaceArea + smallestSide
}

func RibbonNeeded(dim string) int {
	dims := strings.Split(dim, "x")
	l, w, h := toInt(dims[0]), toInt(dims[1]), toInt(dims[2])
	sorted := []int{l, w, h}
	sort.Ints(sorted)
	smallestPerimeter := (2 * sorted[0]) + (2 * sorted[1])
	return smallestPerimeter + (l * w * h)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}
