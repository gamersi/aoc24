package day23p1

import (
	"io"
	"sort"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	network := make(map[string]map[string]bool)

	for _, line := range lines {
		parts := strings.Split(line, "-")
		a, b := parts[0], parts[1]
		if _, ok := network[a]; !ok {
			network[a] = make(map[string]bool)
		}
		network[a][b] = true
		if _, ok := network[b]; !ok {
			network[b] = make(map[string]bool)
		}
		network[b][a] = true
	}

	triples := make(map[[3]string]bool)
	for c1, cs := range network {
		for c2 := range cs {
			for c3 := range network[c2] {
				if c3 == c1 || !network[c3][c1] {
					continue
				}
				triple := [3]string{c1, c2, c3}

				if c1[0] == 't' || c2[0] == 't' || c3[0] == 't' {
					sort.Strings(triple[:])
					triples[triple] = true
				}
			}
		}
	}

	return len(triples)
}
