package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 58, PaperNeeded("2x3x4"))
	assert.Equal(t, 1606483, Part1("input.txt"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 34, RibbonNeeded("2x3x4"))
	assert.Equal(t, 14, RibbonNeeded("1x1x10"))
}
