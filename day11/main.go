package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("1321131112", 40))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part1("1321131112", 50))
}

func Part1(inp string, iter int) int {
	for i := 0; i < iter; i++ {
		inp = lss(inp)
	}
	return len(inp)
}

func lss(s string) (r string) {
	c, times := s[0], 1
	for i := 1; i < len(s); i++ {
		d := s[i]
		if d == c {
			times += 1
			continue
		}
		r += strconv.Itoa(times) + string(c)
		c = d
		times = 1
	}
	return r + strconv.Itoa(times) + string(c)
}