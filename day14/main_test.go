package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 1120, Part1("test1.txt", 1000))
}
func TestPart2(t *testing.T) {
	assert.Equal(t, 688, Part2("test1.txt", 1000))
	assert.Equal(t, 1084, Part2("input.txt", 2503))
}
