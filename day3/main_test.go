package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 2, Part1("test1.txt"))
	assert.Equal(t, 4, Part1("test2.txt"))
	// assert.Equal(t, 58, PaperNeeded("2x3x4"))
	// got := Part1("input.txt")
	// assert.Equal(t, 1606483, got)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 11, Part2("test1.txt"))
	assert.Equal(t, 3, Part2("test2.txt"))
	// got := Part2("input.txt")
	// assert.Equal(t, 1771, got)
}
