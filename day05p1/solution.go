package day05p1

import (
	"io"
	"slices"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	var order_rules [][]int
	var comparator = func(a, b int) int {
		for _, rule := range order_rules {
			if rule[0] == b && rule[1] == a {
				return 1
			}
			if rule[0] == a && rule[1] == b {
				return -1
			}
		}
		return 0
	}

	ruleSection := true

	sum := 0

	for _, line := range lines {
		if line == "" {
			ruleSection = false
			continue
		} else if ruleSection {
			splitNumbers := strings.Split(line, "|")
			numbers := make([]int, len(splitNumbers))
			for i, number := range splitNumbers {
				num, _ := strconv.Atoi(number)
				numbers[i] = num
			}
			order_rules = append(order_rules, numbers)
		} else {
			numbersUpdate := strings.Split(line, ",")
			numbers := make([]int, len(numbersUpdate))
			for i, number := range numbersUpdate {
				num, _ := strconv.Atoi(number)
				numbers[i] = num
			}
			if slices.IsSortedFunc(numbers, comparator) {
				sum += numbers[len(numbers)/2]
			}
		}
	}

	return sum
}

func checkRule(numbers []int, ruleNumbers []int) bool {
	for i, number := range numbers {
		if number == ruleNumbers[0] {
			if i > 0 && numbers[i-1] == ruleNumbers[1] {
				return true
			}
			if i < len(numbers)-1 && numbers[i+1] == ruleNumbers[1] {
				return true
			}
		}
	}
	return false
}
