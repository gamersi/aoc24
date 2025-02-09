package day11p1

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	stones := make(map[int]int)
	for _, line := range strings.Split(lines[0], " ") {
		stone, _ := strconv.Atoi(line)
		stones[stone] = 1
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	total := 0
	for _, stone := range stones {
		total += stone
	}

	return total
}

func blink(stones map[int]int) map[int]int {
	result := make(map[int]int)
	for stone, cnt := range stones {
		if stone == 0 {
			result[1] += cnt
		} else {
			stonestr := strconv.Itoa(stone)
			if len(stonestr)%2 == 0 {
				left, _ := strconv.Atoi(stonestr[:len(stonestr)/2])
				right, _ := strconv.Atoi(stonestr[len(stonestr)/2:])
				result[left] += cnt
				result[right] += cnt
			} else {
				result[stone*2024] += cnt
			}
		}
	}
	return result
}
