package day04p2

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	grid := make([][]rune, len(lines))

	for i, line := range lines {
		for _, c := range line {
			grid[i] = append(grid[i], c)
		}
	}

	height := len(grid)
	width := len(grid[0])

	sum := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'A' {
				if crossMAS(grid, x, y) {
					sum++
				}
			}
		}
	}

	return sum
}

func charAt(grid [][]rune, x, y int) rune {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return 0
	}

	return grid[y][x]
}

func crossMAS(grid [][]rune, x, y int) bool {
	return charAt(grid, x, y) == 'A' &&
		(charAt(grid, x-1, y-1) == 'M' && charAt(grid, x+1, y-1) == 'M' && charAt(grid, x-1, y+1) == 'S' && charAt(grid, x+1, y+1) == 'S' ||
			charAt(grid, x-1, y+1) == 'M' && charAt(grid, x+1, y+1) == 'M' && charAt(grid, x-1, y-1) == 'S' && charAt(grid, x+1, y-1) == 'S' ||
			charAt(grid, x+1, y-1) == 'M' && charAt(grid, x+1, y+1) == 'M' && charAt(grid, x-1, y-1) == 'S' && charAt(grid, x-1, y+1) == 'S' ||
			charAt(grid, x-1, y-1) == 'M' && charAt(grid, x-1, y+1) == 'M' && charAt(grid, x+1, y-1) == 'S' && charAt(grid, x+1, y+1) == 'S')
}
