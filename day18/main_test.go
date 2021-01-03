package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 4, Part1("test1.txt", 4))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 17, Part2("test1.txt", 5))
}
