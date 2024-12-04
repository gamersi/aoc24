package day04p1

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

	const WORD_TO_FIND = "XMAS"

	sum := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'X' {
				sum += countWord(grid, x, y, width, height, WORD_TO_FIND)
			}
		}
	}

	return sum
}

func countWord(grid [][]rune, x, y, width, height int, word string) int {
	directions := [][2]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, 0},
		{0, -1},
		{-1, -1},
	}

	occurrences := 0

	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		found := true
		for i := 1; i < len(word); i++ {
			newX := x + i*dx
			newY := y + i*dy
			if newX < 0 || newX >= width || newY < 0 || newY >= height || grid[newY][newX] != rune(word[i]) {
				found = false
				break
			}
		}
		if found {
			occurrences++
		}
	}
	return occurrences
}
