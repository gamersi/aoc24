package day16p1

import (
	"io"

	"aoc/utils"
)

const TURN_POINTS = 1000
const NORMAL_POINTS = 1

type Dir int

const (
	North Dir = iota
	East
	South
	West
)

var Directions = []utils.Point{
	North: {X: 0, Y: -1},
	East:  {X: 1, Y: 0},
	South: {X: 0, Y: 1},
	West:  {X: -1, Y: 0},
}

type PointDir struct {
	Point utils.Point
	Dir   Dir
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	maze := make([][]rune, len(lines))
	startPos := PointDir{}
	endPos := utils.Point{}

	for i, line := range lines {
		maze[i] = make([]rune, len(line))
		for j, char := range line {
			maze[i][j] = char
			if char == 'S' {
				startPos = PointDir{Point: utils.Point{X: j, Y: i}, Dir: East}
			} else if char == 'E' {
				endPos = utils.Point{X: j, Y: i}
			}
		}
	}

	// Who doesn't love a good BFS

	queue := map[PointDir]int{startPos: 0}
	minScores := map[PointDir]int{startPos: 0}
	current := map[PointDir]int{}

	for len(queue) > 0 {
		current, queue = queue, current
		clear(queue)

		for pd, score := range current {
			for dirIdx, dirPos := range Directions {
				dir := Dir(dirIdx)
				newPos := PointDir{Point: pd.Point.Add(dirPos), Dir: dir}
				if newPos.Point.X < 0 || newPos.Point.X >= len(maze[0]) || newPos.Point.Y < 0 || newPos.Point.Y >= len(maze) {
					continue
				}

				if maze[newPos.Point.Y][newPos.Point.X] == '#' {
					continue
				}

				newScore := score + NORMAL_POINTS
				if dir != pd.Dir {
					newScore += TURN_POINTS
				}
				if minScore, ok := minScores[newPos]; !ok || newScore < minScore {
					minScores[newPos] = newScore
					queue[newPos] = newScore
				}
			}
		}
	}
	return getMinScore(minScores, endPos)
}

func getMinScore(minScores map[PointDir]int, endPos utils.Point) int {
	minScore := 0
	for pd, score := range minScores {
		if pd.Point == endPos {
			if minScore == 0 || score < minScore {
				minScore = score
			}
		}
	}
	return minScore
}
