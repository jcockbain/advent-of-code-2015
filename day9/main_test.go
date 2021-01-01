package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 605, Part1("test1.txt"))
	assert.Equal(t, 117, Part1("input.txt"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 982, Part2("test1.txt"))
	assert.Equal(t, 909, Part2("input.txt"))
}
