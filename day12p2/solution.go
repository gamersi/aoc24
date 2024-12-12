package day12p2

import (
	"io"
	"slices"

	"aoc/utils"
)

type position struct {
	x, y int
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	board := make([][]rune, len(lines))
	for i, line := range lines {
		board[i] = []rune(line)
	}

	sum := 0
	checked := make(map[position]bool)

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			pos := position{x, y}
			if checked[pos] {
				continue
			}
			numChecked := len(checked)
			sides := makeRegion(board, pos, checked)
			area := len(checked) - numChecked

			sum += area * sides
		}
	}

	return sum
}

func makeRegion(board [][]rune, pos position, checked map[position]bool) int {
	sides := 0
	region := make(map[position]bool)
	rtype := board[pos.y][pos.x]
	expandArea(board, pos, rtype, checked, region)
	outDirections := make(map[position][]position)
	for row := range board {
		for col := range board[row] {
			if !region[position{col, row}] {
				continue
			}
			for _, dir := range []position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				newPos := position{col + dir.x, row + dir.y}
				if !oob(board, newPos) && board[newPos.y][newPos.x] == rtype {
					continue
				}
				leftDir := position{dir.y, -dir.x}
				left := position{col + leftDir.x, row + leftDir.y}
				rightDir := position{-dir.y, dir.x}
				right := position{col + rightDir.x, row + rightDir.y}
				if !slices.Contains(outDirections[left], dir) && !slices.Contains(outDirections[right], dir) {
					sides++
				}
				outDirections[position{col, row}] = append(outDirections[position{col, row}], dir)
			}
		}
	}
	return sides
}

func expandArea(board [][]rune, pos position, rtype rune, checked map[position]bool, region map[position]bool) {
	if oob(board, pos) || board[pos.y][pos.x] != rtype {
		return
	}
	if checked[pos] {
		return
	}
	region[pos] = true
	checked[pos] = true
	for _, dir := range []position{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		expandArea(board, position{pos.x + dir.x, pos.y + dir.y}, rtype, checked, region)
	}
}

func oob(board [][]rune, pos position) bool {
	return pos.y < 0 || pos.y >= len(board) || pos.x < 0 || pos.x >= len(board[pos.y])
}
