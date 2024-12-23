package day23p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer string
	}{
		{testInput, "co,de,ka,ta"},
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
