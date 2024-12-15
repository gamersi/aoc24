package day15p1

import (
	"io"

	"aoc/utils"
)

var directions = map[rune]utils.Point{
	'v': {X: 0, Y: 1},
	'^': {X: 0, Y: -1},
	'>': {X: 1, Y: 0},
	'<': {X: -1, Y: 0},
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	linesUntilNewline := 0
	for i, line := range lines {
		if line == "" {
			linesUntilNewline = i
			break
		}
	}
	board := make([][]rune, linesUntilNewline)
	boardMode := true
	moveString := ""

	room := make(map[utils.Point]rune)
	robotPos := utils.Point{X: -1, Y: -1}

	for i, line := range lines {
		if line == "" {
			boardMode = false
			continue
		}
		if boardMode {
			board[i] = []rune(line)
			for j, r := range board[i] {
				room[utils.Point{X: j, Y: i}] = r
				if r == '@' {
					robotPos.X = j
					robotPos.Y = i
				}
			}
		} else {
			moveString += line
		}
	}

	for _, instruction := range moveString {
		dir := directions[instruction]
		robotPos, _ = move(robotPos, dir, room)
	}

	sum := 0
	for p, r := range room {
		if r == 'O' {
			sum += p.X + p.Y*100
		}
	}

	return sum
}

func move(pos utils.Point, dir utils.Point, room map[utils.Point]rune) (utils.Point, bool) {
	newPos := pos.Add(dir)
	switch room[newPos] {
	case '#':
		return pos, false
	case 'O':
		_, possible := move(newPos, dir, room)
		if !possible {
			return pos, false
		}
	}
	room[pos], room[newPos] = '.', room[pos]
	return newPos, true
}
