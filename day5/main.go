package main

import (
	input "aoc2015/inpututils"
	"strings"

	"fmt"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("input.txt"))
}

func Part1(filename string) int {
	lines := input.ReadLines(filename)
	res := 0
	for _, line := range lines {
		if isNice(line) {
			res += 1
		}
	}
	return res
}

func Part2(filename string) int {
	lines := input.ReadLines(filename)
	res := 0
	for _, line := range lines {
		if isNice2(line) {
			res += 1
		}
	}
	return res

}

func isNice(s string) bool {
	if countVowels(s) < 3 {
		return false
	}

	if !hasConsecutiveLetters(s) {
		return false
	}

	if containsSubstrings(s, []string{"ab", "cd", "pq", "xy"}) {
		return false
	}
	return true
}

func isNice2(s string) bool {
	if !containsNonOverlappingPair(s) {
		return false
	}

	if !isTwoAhead(s) {
		return false
	}
	return true
}

func containsNonOverlappingPair(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		pair := string(s[i]) + string(s[i+1])
		if strings.Contains(s[:i], pair) || strings.Contains(s[i+2:], pair) {
			return true
		}
	}
	return false
}

func isTwoAhead(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func countVowels(s string) int {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	count := 0
	for _, c := range s {
		for _, v := range vowels {
			if v == c {
				count += 1
			}
		}
	}
	return count
}

func hasConsecutiveLetters(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func containsSubstrings(s string, substrings []string) bool {
	for _, ss := range substrings {
		if strings.Contains(s, ss) {
			return true
		}
	}
	return false
}
