package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JoinStr(t *testing.T) {

	// input1 := "OneString"
	// input2 := "TwoString"
	// expected := "OneStringTwoString"

	// actual := JoinStr(input1, input2)

	// assert.Equal(t, expected, actual)

	tests := []struct {
		name     string
		input1   string
		input2   string
		expected string
	}{
		{"corrected", "OneString", "TwoString", "OneStringTwoString"},
		{"corrected", "abv", "Asda", "abvAsda"},
		{"corrected", "BABA", "PAPA", "BABAPAPA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := JoinStr(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, result)
		})
	}

}

func Test_SummInteger(t *testing.T) {

	tests := []struct {
		name     string
		input1   int
		input2   int
		expected int
	}{
		{"corrected", -111, 111, 0},
		{"corrected", 0, 0, 0},
		{"corrected", 111, 111, 222},
		{"corrected", 333, 333, 666},
		{"corrected", 999_999_999, 999_999, 1_000_999_998},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SummInteger(tt.input1, tt.input2)
			assert.Equal(t, tt.expected, result)
		})
	}

}
