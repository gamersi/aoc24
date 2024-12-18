package day18p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer string
	}{
		{testInput, "6,1"},
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
