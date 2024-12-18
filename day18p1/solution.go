package day18p1

import (
	"fmt"
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	isTest := len(lines) < 30

	numSimulated := 12
	if !isTest {
		numSimulated = 1024
	}
	size := 6
	if !isTest {
		size = 70
	}

	corrupted := make(map[utils.Point]bool)
	bytes := make([]utils.Point, 0)
	for _, line := range lines {
		x, y := 0, 0
		fmt.Sscanf(line, "%d,%d", &x, &y)
		bytes = append(bytes, utils.Point{X: x, Y: y})
	}
	for i := 0; i < numSimulated; i++ {
		corrupted[bytes[i]] = true
	}

	return steps(corrupted, size)
}

func steps(corrupted map[utils.Point]bool, size int) int {
	type state struct {
		pos   utils.Point
		steps int
	}
	queue := make([]state, 0)
	visited := make(map[utils.Point]bool)
	end := utils.Point{X: size, Y: size}
	var enqueue = func(p utils.Point, s int) {
		for i := 0; i < len(queue); i++ {
			if queue[i].steps > s {
				queue = append(queue[:i], append([]state{{pos: p, steps: s}}, queue[i:]...)...)
				return
			}
		}
		queue = append(queue, state{pos: p, steps: s})
	}
	var dequeue = func() state {
		current := queue[0]
		queue = queue[1:]
		return current
	}
	enqueue(utils.Point{X: 0, Y: 0}, 0)
	for len(queue) > 0 {
		current := dequeue()
		if current.pos == end {
			return current.steps
		}
		if visited[current.pos] {
			continue
		}
		visited[current.pos] = true
		for _, dir := range utils.Directions {
			next := current.pos.Add(dir)
			if next.X < 0 || next.Y < 0 || next.X > size || next.Y > size {
				continue
			}
			if corrupted[next] {
				continue
			}
			enqueue(next, current.steps+1)
		}
	}
	return -1
}
