package day08p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 34},
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
