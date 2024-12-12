package day12p1

import (
	"io"

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
			perimeter := makeRegion(board, pos, checked)
			area := len(checked) - numChecked

			sum += area * perimeter
		}
	}

	return sum
}

func makeRegion(board [][]rune, pos position, checked map[position]bool) int {
	perimeter := 0
	rtype := board[pos.y][pos.x]
	expandArea(board, pos, rtype, checked, &perimeter)
	return perimeter
}

func expandArea(board [][]rune, pos position, rtype rune, checked map[position]bool, perimeter *int) {
	if oob(board, pos) || board[pos.y][pos.x] != rtype {
		*perimeter++
		return
	}
	if checked[pos] {
		return
	}
	checked[pos] = true
	for _, dir := range []position{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		expandArea(board, position{pos.x + dir.x, pos.y + dir.y}, rtype, checked, perimeter)
	}
}

func oob(board [][]rune, pos position) bool {
	return pos.y < 0 || pos.y >= len(board) || pos.x < 0 || pos.x >= len(board[pos.y])
}
