package main

import (
	input "aoc2015/inpututils"
	"fmt"

	"regexp"
	"strconv"
)

var (
	re = regexp.MustCompile(`(.*): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`)
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type ingredientsList []*ingredient

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

func Part1(filename string) int {
	lines := input.ReadLines(filename)
	ingredients := createIngredientsList(lines)
	maximumScore := calculateMaxScore(ingredients)
	return maximumScore
}

func Part2(filename string) int {
	lines := input.ReadLines(filename)
	ingredients := createIngredientsList(lines)
	maximumScore := calculateMaxScoreWithCalLimit(ingredients)
	return maximumScore
}

func createIngredientsList(lines []string) ingredientsList {
	ingredients := make(ingredientsList, len(lines))
	for i, line := range lines {
		parts := re.FindStringSubmatch(line)
		name, cap, dur, fla, tex, cal := parts[1], toInt(parts[2]), toInt(parts[3]), toInt(parts[4]), toInt(parts[5]), toInt(parts[6])
		ingredients[i] = &ingredient{
			name,
			cap,
			dur,
			fla,
			tex,
			cal,
		}
	}
	return ingredients
}

func calculateMaxScore(ingredients ingredientsList) int {
	maxScore := 0

	processSolution := func(a []int, ingredients ingredientsList) {
		if score, _ := calculateScore(a, ingredients); score > maxScore {
			maxScore = score
		}
	}

	buckets := fillBuckets(100, len(ingredients))

	for _, b := range buckets {
		processSolution(b, ingredients)
	}

	return maxScore
}

func calculateMaxScoreWithCalLimit(ingredients ingredientsList) int {
	maxScore := 0

	processSolution := func(a []int, ingredients ingredientsList) {
		if score, calories := calculateScore(a, ingredients); score > maxScore && calories == 500 {
			maxScore = score
		}
	}

	buckets := fillBuckets(100, len(ingredients))

	for _, b := range buckets {
		processSolution(b, ingredients)
	}

	return maxScore
}

func fillBuckets(target int, length int) [][]int {
	if length == 1 {
		return [][]int{[]int{target}}
	}
	res := [][]int{}
	for i := 0; i <= target; i++ {
		rest := fillBuckets(target-i, length-1)
		for _, r := range rest {
			res = append(res, append([]int{i}, r...))
		}
	}
	return res
}

func calculateScore(a []int, ingredients ingredientsList) (int, int) {
	tCap, tDur, tFl, tTex, tCal := 0, 0, 0, 0, 0
	for i := 0; i < len(ingredients); i++ {
		ing := *(ingredients[i])
		num := a[i]
		tCap += (num * ing.capacity)
		tDur += (num * ing.durability)
		tFl += (num * ing.flavor)
		tTex += (num * ing.texture)
		tCal += (num * ing.calories)
	}
	if tCap < 0 || tDur < 0 || tFl < 0 || tTex < 0 {
		return 0, 0
	}
	return tCap * tDur * tFl * tTex, tCal
}

func toInt(s string) int {
	if s[0] == '-' {
		return -1 * toInt(s[1:])
	}
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
