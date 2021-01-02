package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 72, Part1("test1.txt", "d"))
	assert.Equal(t, 507, Part1("test1.txt", "e"))
	assert.Equal(t, 492, Part1("test1.txt", "f"))
	assert.Equal(t, 114, Part1("test1.txt", "g"))
	assert.Equal(t, 65412, Part1("test1.txt", "h"))
	assert.Equal(t, 65079, Part1("test1.txt", "i"))
}
