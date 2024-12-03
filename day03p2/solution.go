package day03p2

import (
	"io"
	"regexp"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	sum := 0

	cumulStr := ""
	for _, line := range lines {
		cumulStr += line
	}

	instructionRegex, _ := regexp.Compile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	matches := instructionRegex.FindAllStringSubmatch(cumulStr, -1)

	mulEnabled := true

	for _, match := range matches {
		instruction := match[1]

		if instruction == "do()" {
			mulEnabled = true
		} else if instruction == "don't()" {
			mulEnabled = false
		} else if match[2] != "" && match[3] != "" {
			if mulEnabled {
				x, _ := strconv.Atoi(match[2])
				y, _ := strconv.Atoi(match[3])
				sum += x * y
			}
		}
	}

	return sum
}

func ArrayIsEmpty(a []int) bool {
	return len(a) == 0
}
