package day09p1

import (
	"io"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	line := lines[0]

	return chksum(defrag(line))
}

func defrag(input string) []int {
	var result []int

	id := 0

	for i, c := range input {
		convInt, _ := strconv.Atoi(string(c))
		storageblock := i % 2 == 0
		for j := 0; j < convInt; j++ {
			if storageblock {
				result = append(result, id)
			} else {
				result = append(result, -1)
			}
		}
		if storageblock {
			id++
		}
	}

	for i := len(result)-1; i > 0; i-- {
		if result[i] != -1 {
			firstEmpty := -1
			for j := 0; j < len(result); j++ {
				if result[j] == -1 {
					firstEmpty = j
					break
				}
			}
			if firstEmpty != -1 && firstEmpty < i {
				result[firstEmpty] = result[i]
				result[i] = -1
			} else {
				break
			}
		}
	}

	return result
}

func chksum(input []int) int {
	sum := 0

	for i, id := range input {
		if id != -1 {
			sum += i * id
		}
	}

	return sum
}
