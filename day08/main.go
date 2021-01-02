package main

import (
	input "aoc2015/inpututils"

	"fmt"
	"regexp"
	"strings"
)

var (
	hexRe   = regexp.MustCompile(`\\x[\d{2}|[a-z]{2}`)
	quoteRe = regexp.MustCompile(`\\"`)
	backRe  = regexp.MustCompile(`\\\\`)

	backRe2  = regexp.MustCompile(`\\`)
	quoteRe2 = regexp.MustCompile(`\"`)
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
		escapeString := newEscapeString(line)
		total += len(escapeString.rawString) - len(escapeString.escapedString)
	}
	return total
}

func Part2(filename string) int {
	lines := input.ReadLines(filename)
	total := 0

	for _, line := range lines {
		escapeString := newEscapeString2(line)
		total += len(escapeString.escapedString) - len(escapeString.rawString)
	}
	return total
}

type EscapeString struct {
	rawString     string
	escapedString string
}

func newEscapeString(s string) EscapeString {
	// use X as placeholder for escaped char
	escaped := strings.TrimPrefix(s, `"`)
	escaped = strings.TrimSuffix(escaped, `"`)
	escaped = backRe.ReplaceAllString(escaped, `X`)
	escaped = quoteRe.ReplaceAllString(escaped, `X`)
	escaped = hexRe.ReplaceAllString(escaped, `X`)
	return EscapeString{
		s,
		escaped,
	}

}

func newEscapeString2(s string) EscapeString {
	// use X as placeholder for escaped char
	escaped := backRe2.ReplaceAllString(s, `XX`)
	escaped = quoteRe2.ReplaceAllString(escaped, `XX`)
	return EscapeString{
		s,
		`"` + escaped + `"`,
	}

}
