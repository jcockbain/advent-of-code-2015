package main

import (
	"fmt"

	"regexp"
)

var (
	forbiddenRe = regexp.MustCompile(`i|l|o`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("hxbxwxba"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part1("hxbxxzaa"))

}

func Part1(s string) string {
	for !isValid(s) {
		s = nextPword(s)
	}
	return s
}

func nextPword(s string) string {
	sArray := []byte(s)
	for i := len(s) - 1; i >= 0; i-- {
		c := sArray[i]
		if c != 'z' {
			sArray[i] = byte(int(c) + 1)
			break
		} else {
			sArray[i] = 'a'
		}
	}
	return string(sArray)
}

func isValid(s string) bool {
	return !containsForbiddenLetters(s) && containsRepeatingLetters(s) && hasConsecutiveLetters(s)
}

func containsForbiddenLetters(s string) bool {
	return forbiddenRe.MatchString(s)
}

func containsRepeatingLetters(s string) bool {
	repeatingLetters := 0
	i := 0
	for i < len(s)-1 {
		if s[i] == s[i+1] {
			repeatingLetters += 1
			i += 2
		} else {
			i += 1
		}
	}
	return repeatingLetters >= 2
}

func hasConsecutiveLetters(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		d1 := int(s[i+1]) - int(s[i])
		d2 := int(s[i+2]) - int(s[i+1])
		if (d1 == 1) && (d2 == 1) {
			return true
		}
	}
	return false
}
