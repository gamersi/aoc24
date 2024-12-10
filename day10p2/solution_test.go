package day10p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 81},
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
