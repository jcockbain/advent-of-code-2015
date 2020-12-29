package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 998996, Part1("test1.txt"))
}
func TestPart2(t *testing.T) {
	assert.Equal(t, 15343601, Part2("input.txt"))
}
