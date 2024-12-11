package day11p2

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	stones := make(map[int64]int64)
	for _, line := range strings.Split(lines[0], " ") {
		stone, _ := strconv.ParseInt(line, 10, 64)
		stones[stone] = 1
	}

	for i := 0; i < 75; i++ {
		stones = blink(stones)
	}

	total := int64(0)
	for _, stone := range stones {
		total += stone
	}

	return total
}

func blink(stones map[int64]int64) map[int64]int64 {
	result := make(map[int64]int64)
	for stone, cnt := range stones {
		if stone == 0 {
			result[1] += cnt
		} else {
			stonestr := strconv.FormatInt(stone, 10)
			if len(stonestr)%2 == 0 {
				left, _ := strconv.Atoi(stonestr[:len(stonestr)/2])
				right, _ := strconv.Atoi(stonestr[len(stonestr)/2:])
				result[int64(left)] += cnt
				result[int64(right)] += cnt
			} else {
				result[stone*2024] += cnt
			}
		}
	}
	return result
}
