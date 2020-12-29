package main

import (
	input "aoc2015/inpututils"

	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

type pos struct {
	x, y int
}

func Part1(filename string) int {
	lines := input.ReadLines(filename)
	lights := map[pos]bool{}

	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			p := pos{x, y}
			lights[p] = false
		}
	}

	re := regexp.MustCompile(`(.+) (\d+),(\d+) through (\d+),(\d+)`)

	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		action := parts[1]
		lowerStart, lowerEnd := toInt(parts[2]), toInt(parts[3])
		upperStart, upperEnd := toInt(parts[4]), toInt(parts[5])

		for y := lowerEnd; y <= upperEnd; y++ {
			for x := lowerStart; x <= upperStart; x++ {
				p := pos{x, y}
				switch action {
				case "turn on":
					lights[p] = true
				case "turn off":
					lights[p] = false
				case "toggle":
					lights[p] = !lights[p]
				}
			}
		}
	}

	res := 0
	for _, on := range lights {
		if on {
			res += 1
		}
	}
	return res
}

func Part2(filename string) int {
	lines := input.ReadLines(filename)
	lights := map[pos]int{}

	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			p := pos{x, y}
			lights[p] = 0
		}
	}

	re := regexp.MustCompile(`(.+) (\d+),(\d+) through (\d+),(\d+)`)

	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		action := parts[1]
		lowerStart, lowerEnd := toInt(parts[2]), toInt(parts[3])
		upperStart, upperEnd := toInt(parts[4]), toInt(parts[5])

		for y := lowerEnd; y <= upperEnd; y++ {
			for x := lowerStart; x <= upperStart; x++ {
				p := pos{x, y}
				switch action {
				case "turn on":
					lights[p] += 1
				case "turn off":
					if lights[p] != 0 {
						lights[p] -= 1
					}
				case "toggle":
					lights[p] += 2
				}
			}
		}
	}

	res := 0
	for _, brightness := range lights {
		res += brightness
	}
	return res
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
