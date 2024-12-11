package day11p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `125 17`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int64
	}{
		{testInput, int64(65601038650482)},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := Solve(r).(int64)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
