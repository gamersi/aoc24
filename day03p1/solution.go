package day03p1

import (
	"io"
	"regexp"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	sum := 0

	instructionRegex, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)

	for _, line := range lines {
		matches := instructionRegex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			sum += StrToInt(match[1]) * StrToInt(match[2])
		}
	}

	return sum
}

func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
