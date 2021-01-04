package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1(36000000))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2(36000000))

}

func Part1(minPresents int) int {
	for i := 1; i < minPresents; i++ {
		if sumFactors(i) >= minPresents/10 {
			return i
		}
	}
	return -1
}

func Part2(minPresents int) int {
	for i := 1; i < minPresents; i++ {
		if sumFactorsLessThan50(i) >= minPresents/11 {
			return i
		}
	}
	return -1
}

func sumFactors(i int) int {
	s := 0
	f := getFactors(i)
	for _, i := range f {
		s += i
	}
	return s
}

func getFactors(x int) []int {
	factors := []int{}
	for i := 1; i <= int(math.Floor(math.Sqrt(float64(x)))); i++ {
		if x%i == 0 {
			factors = append(factors, i)
			factors = append(factors, x/i)
		}
	}
	return factors
}

func sumFactorsLessThan50(i int) int {
	s := 0
	f := getFactorsLessThan50(i)
	for _, i := range f {
		s += i
	}
	return s
}

func getFactorsLessThan50(x int) []int {
	factors := []int{}
	for i := 1; i <= 50; i++ {
		if x%i == 0 {
			factors = append(factors, i)
			factors = append(factors, x/i)
		}
	}
	return factors
}
