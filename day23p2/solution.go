package day23p2

import (
	"io"
	"sort"
	"strings"

	"aoc/utils"
)

type Party map[string]bool

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

	var largestParty Party

	for c1 := range network {
		party := dfs(network, c1, make(Party))
		if len(party) > len(largestParty) {
			largestParty = party
		}
	}
	largestPartyPW := make([]string, 0, len(largestParty))
	for k := range largestParty {
		largestPartyPW = append(largestPartyPW, k)
	}
	sort.Strings(largestPartyPW)

	return strings.Join(largestPartyPW, ",")
}

func dfs(network map[string]map[string]bool, c string, visited Party) Party {
	visited[c] = true
	var largestParty Party
NextParty:
	for c2 := range network[c] {
		if visited[c2] {
			continue
		}
		for prev := range visited {
			if !network[c2][prev] {
				continue NextParty
			}
		}
		party := dfs(network, c2, visited)
		if len(party) > len(largestParty) {
			largestParty = party
		}
	}
	if len(largestParty) == 0 {
		// deep clone visited
		largestParty = make(Party)
		for k, v := range visited {
			largestParty[k] = v
		}
	}
	return visited
}
