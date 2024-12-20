package day20p1

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	track := make([][]rune, len(lines))
	startPos := utils.Point{}
	endPos := utils.Point{}
	for i, line := range lines {
		for j, c := range line {
			if c == 'S' {
				startPos = utils.Point{X: j, Y: i}
			} else if c == 'E' {
				endPos = utils.Point{X: j, Y: i}
			}
		}
		track[i] = []rune(line)
	}

	stepsFwd, fwdVisited := bfs(track, startPos, endPos)
	stepsBwd, bwdVisited := bfs(track, endPos, startPos)
	if stepsFwd == -1 || stepsBwd == -1 || stepsFwd != stepsBwd {
		return -1
	}

	nCheats := 0

	const MAX_CHEAT_TIME = 2
	const SAVE_MIN = 100

	for pos, step := range fwdVisited {
		if step > stepsFwd-SAVE_MIN {
			continue
		}
		for dy := -MAX_CHEAT_TIME; dy <= MAX_CHEAT_TIME; dy++ {
			timeLeft := MAX_CHEAT_TIME - utils.Abs(dy)
			for dx := -timeLeft; dx <= timeLeft; dx++ {
				cheatPos := pos.Add(utils.Point{X: dx, Y: dy})
				if cheatPos.X < 0 || cheatPos.X >= len(track[0]) || cheatPos.Y < 0 || cheatPos.Y >= len(track) {
					continue
				}
				if track[cheatPos.Y][cheatPos.X] == '#' {
					continue
				}
				time := utils.Abs(dy) + utils.Abs(dx)
				if bwdStep, ok := bwdVisited[cheatPos]; ok {
					if step+time+bwdStep <= stepsFwd-SAVE_MIN {
						nCheats++
					}
				}
			}
		}
	}

	return nCheats

}

func bfs(track [][]rune, startPos, endPos utils.Point) (int, map[utils.Point]int) {
	visited := map[utils.Point]int{startPos: 0}
	queue := []utils.Point{startPos}
	var curr []utils.Point

	step := 0
	for len(queue) > 0 {
		step++
		curr, queue = queue, curr[:0]
		for _, pos := range curr {
			if pos == endPos {
				return step - 1, visited
			}
			for _, dir := range utils.Directions {
				nextPos := pos.Add(dir)
				if nextPos.X < 0 || nextPos.X >= len(track[0]) || nextPos.Y < 0 || nextPos.Y >= len(track) {
					continue
				}
				if track[nextPos.Y][nextPos.X] == '#' {
					continue
				}
				if _, ok := visited[nextPos]; ok {
					continue
				}
				visited[nextPos] = step
				queue = append(queue, nextPos)
			}
		}
	}
	return -1, visited
}
