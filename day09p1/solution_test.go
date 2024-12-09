package day09p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `2333133121414131402`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 1928},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := Solve(r).(int)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
