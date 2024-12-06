package day06p1

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	player := [2]int{0, 0}
	vel := [2]int{0, -1}

	const OBSTACLE = '#'

	board := make([][]rune, len(lines))
	for i, line := range lines {
		for j, c := range line {
			if c == '^' {
				player = [2]int{j, i}
			}
			board[i] = append(board[i], c)
		}
	}


	height := len(board)
	width := len(board[0])

	visited := make([][]bool, height)
	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
	}

	for player[0] + vel[0] >= 0 && player[0] + vel[0] < height && player[1] + vel[1] >= 0 && player[1] + vel[1] < width {
		if board[player[1] + vel[1]][player[0] + vel[0]] == OBSTACLE {
			if vel[0] == 0 && vel[1] == -1 {
				vel = [2]int{1, 0}
			} else if vel[0] == 1 && vel[1] == 0 {
				vel = [2]int{0, 1}
			} else if vel[0] == 0 && vel[1] == 1 {
				vel = [2]int{-1, 0}
			} else if vel[0] == -1 && vel[1] == 0 {
				vel = [2]int{0, -1}
			} else {
				panic("invalid velocity")
			}
		} else {
			player[0] += vel[0]
			player[1] += vel[1]
			visited[player[1]][player[0]] = true
		}
	}

	return countVisited(visited)
}

func countVisited(visited [][]bool) int {
	count := 0
	for _, row := range visited {
		for _, v := range row {
			if v {
				count++
			}
		}
	}
	return count
}
