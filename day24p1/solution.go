package day24p1

import (
	"fmt"
	"io"
	"sort"

	"aoc/utils"
)

type Gate struct {
	Op    string
	Left  string
	Right string
	Dest  string
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	wires := make(map[string]byte)
	gates := make(map[Gate]bool)

	wireDecs := true

	for _, line := range lines {
		if line == "" {
			wireDecs = false
			continue
		} else if wireDecs {
			var wire string
			var val byte
			if _, err := fmt.Sscanf(line, "%3s: %d", &wire, &val); err != nil {
				panic(err)
			}
			wires[wire] = val
		} else {
			var g Gate
			if _, err := fmt.Sscanf(line, "%s %s %s -> %s", &g.Left, &g.Op, &g.Right, &g.Dest); err != nil {
				panic(err)
			}
			gates[g] = false
		}
	}

	for len(gates) > 0 {
		for g := range gates {
			if g.Left != "" {
				if _, ok := wires[g.Left]; !ok {
					continue
				}
			}
			if g.Right != "" {
				if _, ok := wires[g.Right]; !ok {
					continue
				}
			}
			switch g.Op {
			case "AND":
				wires[g.Dest] = wires[g.Left] & wires[g.Right]
			case "OR":
				wires[g.Dest] = wires[g.Left] | wires[g.Right]
			case "XOR":
				wires[g.Dest] = wires[g.Left] ^ wires[g.Right]
			}
			delete(gates, g)
		}
	}

	return produceResult(wires)
}

func produceResult(wires map[string]byte) int {
	var keys []string
	for k := range wires {
		if k[0] == 'z' {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)

	var result int
	for i, k := range keys {
		if wires[k] == 1 {
			result |= 1 << i
		}
	}
	return result
}
