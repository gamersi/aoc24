package day24p2

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"aoc/utils"
)

type Gate struct {
	Op    string
	Left  string
	Right string
	Dest  string
}

func notIn(value byte, list []byte) bool {
	for _, v := range list {
		if v == value {
			return false
		}
	}
	return true
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

	highestZ := "z00"

	for g := range gates {
		if g.Dest[0] == 'z' {
			znum, _ := strconv.Atoi(g.Dest[1:])
			highestZnum, _ := strconv.Atoi(highestZ[1:])
			if znum > highestZnum {
				highestZ = g.Dest
			}
		}
	}

	wrong := make([]string, 0)

	for g := range gates {
		if g.Dest[0] == 'z' && g.Op != "XOR" && g.Dest != highestZ {
			wrong = append(wrong, g.Dest)
		}
		if g.Op == "XOR" &&
			notIn(g.Dest[0], []byte{'x', 'y', 'z'}) &&
			notIn(g.Left[0], []byte{'x', 'y', 'z'}) &&
			notIn(g.Right[0], []byte{'x', 'y', 'z'}) {
			wrong = append(wrong, g.Dest)
		}
		if g.Op == "AND" && (g.Left != "x00" && g.Right != "x00") {
			for g2 := range gates {
				if (g.Dest == g2.Left || g.Dest == g2.Right) && g2.Op != "OR" {
					wrong = append(wrong, g.Dest)
				}
			}
		}
		if g.Op == "XOR" {
			for g2 := range gates {
				if (g.Dest == g2.Left || g.Dest == g2.Right) && g2.Op == "OR" {
					wrong = append(wrong, g.Dest)
				}
			}
		}
	}

	sort.Strings(wrong)

	wrong = utils.RemoveDuplicates(wrong)

	return strings.Join(wrong, ",")
}
