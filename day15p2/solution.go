package day15p2

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
				switch r {
				case '@':
					room[utils.Point{X: j * 2, Y: i}] = r
					room[utils.Point{X: j*2 + 1, Y: i}] = '.'
					robotPos.X = j * 2
					robotPos.Y = i
				case 'O':
					room[utils.Point{X: j * 2, Y: i}] = '['
					room[utils.Point{X: j*2 + 1, Y: i}] = ']'
				default:
					room[utils.Point{X: j * 2, Y: i}] = r
					room[utils.Point{X: j*2 + 1, Y: i}] = r
				}
			}
		} else {
			moveString += line
		}
	}

	for _, instruction := range moveString {
		dir := directions[instruction]
		if movePossible(robotPos, dir, room) {
			robotPos, _ = move(robotPos, dir, room)
		}
	}

	sum := 0
	for p, r := range room {
		if r == '[' {
			sum += p.X + p.Y*100
		}
	}

	return sum
}

func movePossible(pos utils.Point, dir utils.Point, room map[utils.Point]rune) bool {
	newPos := pos.Add(dir)
	switch room[newPos] {
	case '#':
		return false
	case '[':
		if dir.X == 0 {
			return movePossible(newPos, dir, room) && movePossible(newPos.Add(utils.Point{X: 1, Y: 0}), dir, room)
		}
		return movePossible(newPos, dir, room)
	case ']':
		if dir.X == 0 {
			return movePossible(newPos, dir, room) && movePossible(newPos.Add(utils.Point{X: -1, Y: 0}), dir, room)
		}
		return movePossible(newPos, dir, room)
	}
	return true
}

func move(pos utils.Point, dir utils.Point, room map[utils.Point]rune) (utils.Point, bool) {
	newPos := pos.Add(dir)
	switch room[newPos] {
	case '#':
		return pos, false
	case '[':
		if dir.X == 0 {
			move(newPos.Add(utils.Point{X: 1, Y: 0}), dir, room)
		}
		move(newPos, dir, room)
	case ']':
		if dir.X == 0 {
			move(newPos.Add(utils.Point{X: -1, Y: 0}), dir, room)
		}
		move(newPos, dir, room)
	}
	room[pos], room[newPos] = '.', room[pos]
	return newPos, true
}
