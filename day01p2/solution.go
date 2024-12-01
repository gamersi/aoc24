package day01p2

import (
	"io"
	"log"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	sum := 0
	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))

	for i, line := range lines {
		cols := strings.Split(line, "   ")
		if len(cols) != 2 {
			log.Panicf("Invalid input: %s", line)
		}
		num1, err := strconv.Atoi(cols[0])
		if err != nil {
			log.Panicf("Invalid input: %s", line)
		}
		num2, err := strconv.Atoi(cols[1])
		if err != nil {
			log.Panicf("Invalid input: %s", line)
		}
		list1[i] = num1
		list2[i] = num2
	}

	counts := make(map[int]int)
	for _, num := range list2 {
		counts[num]++
	}

	for i := 0; i < len(list1); i++ {
		sum += list1[i] * counts[list1[i]]
	}

	return sum
}
