package day17p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer uint64
	}{
		{testInput, 117440},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := Solve(r).(uint64)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
