package day11p1

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	arrStr := strings.Split(lines[0], " ")
	arr := make([]int, len(arrStr))
	for i, s := range arrStr {
		arr[i], _ = strconv.Atoi(s)
	}

	for i := 0; i < 25; i++ {
		arr = blink(arr)
	}

	return len(arr)
}

func blink(arr []int) []int {
    var result []int
    for _, stone := range arr {
        stonestr := strconv.Itoa(stone)
        if stone == 0 {
            result = append(result, 1)
        } else if len(stonestr)%2 == 0 {
            part1 := stonestr[:len(stonestr)/2]
            part2 := stonestr[len(stonestr)/2:]
            part1int, _ := strconv.Atoi(part1)
            part2int, _ := strconv.Atoi(part2)
            result = append(result, part1int, part2int)
        } else {
            result = append(result, stone*2024)
        }
    }
    return result
}
