package day19p2

import (
	"io"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	availableTowels := strings.Split(lines[0], ", ")

	designs := make([]string, 0)
	for _, line := range lines[2:] {
		designs = append(designs, line)
	}

	return countPossibleDesigns(designs, availableTowels)
}

func countPossibleDesigns(designs, availableTowels []string) int {
	count := 0
	for _, design := range designs {
		count += getPossibleDesigns(design, availableTowels)
	}
	return count
}

var memo = make(map[string]int)

func getPossibleDesigns(design string, availableTowels []string) int {
	if len(design) == 0 {
		return 1
	}
	if v, ok := memo[design]; ok {
		return v
	}
	result := 0
	for _, towel := range availableTowels {
		if len(towel) <= len(design) && towel == design[0:len(towel)] {
			numPossible := getPossibleDesigns(design[len(towel):], availableTowels)
			result += numPossible
			memo[design[len(towel):]] = numPossible
		}
	}
	return result
}
