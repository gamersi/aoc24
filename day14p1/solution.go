package day14p1

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	isTest := len(lines) == 12
	boardWidth := 11
	if !isTest {
		boardWidth = 101
	}
	boardHeight := 7
	if !isTest {
		boardHeight = 103
	}

	endPoints := make(map[utils.Point]int)

	for _, line := range lines {
		ints := utils.GetInts(line)
		px, py, vx, vy := ints[0], ints[1], ints[2], ints[3]
		for range 100 {
			px += vx
			py += vy
			if px < 0 {
				px += boardWidth
			} else if px >= boardWidth {
				px -= boardWidth
			}
			if py < 0 {
				py += boardHeight
			} else if py >= boardHeight {
				py -= boardHeight
			}
		}
		endPoints[utils.Point{X: px, Y: py}]++
	}

	bots := make([]int, 4)

	for point, count := range endPoints {
		if point.X < boardWidth/2 && point.Y < boardHeight/2 {
			bots[0] += count
		} else if point.X > boardWidth/2 && point.Y < boardHeight/2 {
			bots[1] += count
		} else if point.X < boardWidth/2 && point.Y > boardHeight/2 {
			bots[2] += count
		} else if point.X > boardWidth/2 && point.Y > boardHeight/2 {
			bots[3] += count
		}
	}

	return multAll(bots)
}

func multAll(arr []int) int {
	result := 1
	for _, n := range arr {
		result *= n
	}
	return result
}
