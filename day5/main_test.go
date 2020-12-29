package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, true, isNice("ugknbfddgicrmopn"))
	assert.Equal(t, true, isNice("aaa"))
	assert.Equal(t, false, isNice("jchzalrnumimnmhp"))
	assert.Equal(t, false, isNice("haegwjzuvuyypxyu"))
	assert.Equal(t, false, isNice("dvszwmarrgswjxmb"))
	// assert.Equal(t, 58, PaperNeeded("2x3x4"))
	// got := Part1("input.txt")
	// assert.Equal(t, 1606483, got)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, true, isNice2("qjhvhtzxzqqjkmpb"))
	assert.Equal(t, true, isNice2("xxyxx"))
	assert.Equal(t, false, isNice2("uurcxstgmygtbstg"))
	assert.Equal(t, false, isNice2("ieodomkazucvgmuy"))

}
