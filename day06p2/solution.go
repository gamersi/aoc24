package day06p2

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	player := [2]int{0, 0}

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

	numOfLoops := 0
	for y, row := range board {
		for x, c := range row {
			if c != '^' && c != '#' {
				board[y][x] = '#'
				if isLoop(board, player[0], player[1]) {
					numOfLoops++
				}
				board[y][x] = '.'
			}
		}
	}

	return numOfLoops
}

func isLoop(board [][]rune, x, y int) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	width := len(board[0])
	height := len(board)

	steps := 1

	player := [2]int{x, y}
	vel := [2]int{0, -1}

	visited[y][x] = true

	for player[0] + vel[0] >= 0 && player[0] + vel[0] < height && player[1] + vel[1] >= 0 && player[1] + vel[1] < width {
		if steps - width * height >= 1 {
			return true
		}
		if board[player[1] + vel[1]][player[0] + vel[0]] == '#' {
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
			steps++
		}
		visited[y][x] = true
	}

	return false
}
