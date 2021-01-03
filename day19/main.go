package main

import (
	input "aoc2015/inpututils"
	"fmt"

	"regexp"
	"strings"
)

var (
	re = regexp.MustCompile(`(.*) => (.*)`)
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("input.txt"))
}

type replacement struct {
	out string
	in  string
}

func Part1(filename string) int {
	inp := input.ReadRaw(filename)
	split := strings.Split(inp, "\n\n")

	repl, molecule := split[0], split[1]
	splitRepl := strings.Split(repl, "\n")
	replacements := make([]replacement, len(splitRepl))

	for i, line := range splitRepl {
		parts := re.FindStringSubmatch(line)
		replacements[i] = replacement{
			parts[1],
			parts[2],
		}
	}

	seen := map[string]bool{}

	for _, repl := range replacements {
		replRe := regexp.MustCompile(repl.out)
		pos := replRe.FindAllStringIndex(molecule, -1)
		if pos != nil {
			for _, p := range pos {
				newS := molecule[:p[0]] + repl.in + molecule[p[1]:]
				seen[newS] = true
			}
		}
	}
	return len(seen)
}
