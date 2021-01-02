package main

import (
	input "aoc2015/inpututils"
	"fmt"
	"strconv"

	"regexp"
)

var (
	re = regexp.MustCompile(`(.*) can fly (\d+) km\/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt", 2503))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt", 2503))
}

type reindeer struct {
	name        string
	speed       int
	restingTime int
	flyingTime  int
	isFlying    bool
	timeLeft    int
	distance    int
	score       int
}

func Part1(filename string, time int) int {
	lines := input.ReadLines(filename)
	reindeers := []*reindeer{}
	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		name, speed, flyTime, restTime := parts[1], toInt(parts[2]), toInt(parts[3]), toInt(parts[4])
		reindeers = append(reindeers, &reindeer{
			name,
			speed,
			restTime,
			flyTime,
			true,
			flyTime,
			0,
			0,
		})
	}

	for i := 0; i < time; i++ {
		for i, rPointer := range reindeers {
			r := *rPointer
			r.timeLeft -= 1
			if r.isFlying {
				r.distance += r.speed
			}
			if r.timeLeft == 0 {
				r.isFlying = !r.isFlying
				if r.isFlying {
					r.timeLeft = r.flyingTime
				} else {
					r.timeLeft = r.restingTime
				}
			}
			reindeers[i] = &r
		}
	}

	maxDistance := 0
	for _, r := range reindeers {
		if r.distance > maxDistance {
			maxDistance = r.distance
		}
	}
	return maxDistance
}

func Part2(filename string, time int) int {
	lines := input.ReadLines(filename)
	reindeers := []*reindeer{}
	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		name, speed, flyTime, restTime := parts[1], toInt(parts[2]), toInt(parts[3]), toInt(parts[4])
		reindeers = append(reindeers, &reindeer{
			name,
			speed,
			restTime,
			flyTime,
			true,
			flyTime,
			0,
			0,
		})
	}

	for i := 0; i < time; i++ {
		maxDistance, maxIdx := (*reindeers[0]).distance, 0
		for i, rPointer := range reindeers {
			r := *rPointer
			r.timeLeft -= 1
			if r.isFlying {
				r.distance += r.speed
			}
			if r.timeLeft == 0 {
				r.isFlying = !r.isFlying
				if r.isFlying {
					r.timeLeft = r.flyingTime
				} else {
					r.timeLeft = r.restingTime
				}
			}
			reindeers[i] = &r
			if r.distance > maxDistance {
				maxDistance = r.distance
				maxIdx = i
			}
		}
		(*reindeers[maxIdx]).score += 1
	}

	maxScore := 0
	for _, r := range reindeers {
		if r.score > maxScore {
			maxScore = r.score
		}
	}
	return maxScore
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
