package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPword(t *testing.T) {

}

func TestPart1(t *testing.T) {
	assert.Equal(t, false, isValid("hijklmmn"))
	assert.Equal(t, false, isValid("abbceffg"))
	assert.Equal(t, false, isValid("abbcegjk"))
	assert.Equal(t, true, isValid("ghjaabcc"))

	assert.Equal(t, "abcdffab", nextPword("abcdffaa"))
	assert.Equal(t, "abcdffba", nextPword("abcdffaz"))
}
