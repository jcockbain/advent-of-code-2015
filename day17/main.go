package main

import (
	input "aoc2015/inpututils"
	"fmt"

	"math"
	"regexp"
	"strconv"
)

var (
	re = regexp.MustCompile(`[a-z]*: (\d+)`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt", 150))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt", 150))
}

func Part1(filename string, cap int) int {
	lines := input.ReadLines(filename)
	containers := make([]int, len(lines))
	for i, c := range lines {
		containers[i] = toInt(c)
	}
	return countNumberOfWays(containers, cap)
}

func countNumberOfWays(containers []int, target int) int {
	var countWays func(int, int) int

	countWays = func(idx int, target int) int {
		if target == 0 {
			return 1
		}
		if idx == len(containers) || target < 0 {
			return 0
		}
		w1 := countWays(idx+1, target)
		w2 := countWays(idx+1, target-containers[idx])
		return w1 + w2
	}

	return countWays(0, target)
}

func Part2(filename string, cap int) int {
	lines := input.ReadLines(filename)
	containers := make([]int, len(lines))
	for i, c := range lines {
		containers[i] = toInt(c)
	}
	return countNumberOfMinimumWays(containers, cap)
}

func countNumberOfMinimumWays(containers []int, target int) int {
	var countWays func(int, int, int)
	minimumContainers := math.MaxInt32
	minCount := 0

	countWays = func(idx int, target int, containersUsed int) {
		if target == 0 {
			if containersUsed == minimumContainers {
				minCount += 1
			} else if containersUsed < minimumContainers {
				minimumContainers = containersUsed
				minCount = 1
			}
			return
		}
		if idx == len(containers) || target < 0 {
			return
		}
		countWays(idx+1, target, containersUsed)
		countWays(idx+1, target-containers[idx], containersUsed+1)
	}

	countWays(0, target, 0)
	return minCount
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
