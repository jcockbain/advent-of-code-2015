package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 609043, Part1("abcdef"))
	assert.Equal(t, 1048970, Part1("pqrstuv"))
}

// func TestPart2(t *testing.T) {
// 	assert.Equal(t, 11, Part2("test1.txt"))
// 	assert.Equal(t, 3, Part2("test2.txt"))
// 	// got := Part2("input.txt")
// 	// assert.Equal(t, 1771, got)
// }
