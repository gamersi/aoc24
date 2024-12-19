package day19p1

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
		if isPossibleDesign(design, availableTowels) {
			count++
		}
	}
	return count
}

func isPossibleDesign(design string, availableTowels []string) bool {
	if len(design) == 0 {
		return true
	}
	for _, towel := range availableTowels {
		if len(towel) <= len(design) && towel == design[:len(towel)] {
			if isPossibleDesign(design[len(towel):], availableTowels) {
				return true
			}
		}
	}
	return false
}
