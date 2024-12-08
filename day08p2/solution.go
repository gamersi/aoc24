package day08p2

import (
	"io"

	"aoc/utils"
)

type position struct {
	x, y int
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	antennaPositions := make(map[rune][]position)
	antinodes := make(map[position]bool)
	for i, line := range lines {
		for j, c := range line {
			if c == '.' {
				continue
			}
			antennaPositions[c] = append(antennaPositions[c], position{j, i})
		}
	}

	for distance := 0; findAntinodes(antennaPositions, antinodes, distance, len(lines), len(lines[0])) > 0; distance++ {
	}

	return len(antinodes)
}

func findAntinodes(antennaPositions map[rune][]position, antinodes map[position]bool, distance int, rows, cols int) int {
	prevLen := len(antinodes)
	for _, aps := range antennaPositions {
		for i := 0; i < len(aps)-1; i++ {
			for j := i + 1; j < len(aps); j++ {
				antinode(rows, cols, position{aps[i].y - distance*(aps[j].y-aps[i].y), aps[i].x - distance*(aps[j].x-aps[i].x)}, antinodes)
				antinode(rows, cols, position{aps[j].y - distance*(aps[i].y-aps[j].y), aps[j].x - distance*(aps[i].x-aps[j].x)}, antinodes)
			}
		}
	}
	return len(antinodes) - prevLen
}

func antinode(rows, cols int, position position, antinodes map[position]bool) {
	if position.x >= 0 && position.x < rows && position.y >= 0 && position.y < cols {
		antinodes[position] = true
	}
}
