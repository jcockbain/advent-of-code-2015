package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	mockBoss := player{12, 7, 2}
	mockPlayer := player{8, 5, 5}
	assert.Equal(t, true, mockPlayer.beats(mockBoss))
}
