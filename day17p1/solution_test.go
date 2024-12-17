package day17p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer string
	}{
		{testInput, "4,6,3,5,6,3,5,2,1,0"},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := Solve(r).(string)

		if result != test.answer {
			t.Errorf("Expected %s, got %s", test.answer, result)
		}
	}
}
