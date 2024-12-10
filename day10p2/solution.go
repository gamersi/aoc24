package day10p2

import (
	"io"

	"aoc/utils"
)

type position struct {
	x, y int
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	board := make([][]int, len(lines))
	for i, line := range lines {
		board[i] = make([]int, len(line))
		for j, c := range line {
			board[i][j] = int(c) - int('0')
		}
	}

	scores := 0

	for i, row := range board {
		for j, _ := range row {
			if row[j] != 0 {
				continue
			}
			queue := []position{{i, j}}
			for len(queue) > 0 {
				pos := queue[0]
				queue = queue[1:]
				if board[pos.x][pos.y] == 9 {
					scores++
				}
				for _, dir := range []position{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
					if pos.x+dir.x >= 0 && pos.x+dir.x < len(board) && pos.y+dir.y >= 0 && pos.y+dir.y < len(board[0]) && (board[pos.x+dir.x][pos.y+dir.y]-board[pos.x][pos.y]) == 1 {
						queue = append(queue, position{pos.x + dir.x, pos.y + dir.y})
					}
				}
			}
		}
	}

	return scores
}
