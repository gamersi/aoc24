package day02p2

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

	for _, line := range lines {
		nums := strings.Split(line, " ")

		isValid := checkBruteforceNumbers(nums)

		if isValid {
			sum++
		}
	}

	return sum
}

func checkBruteforceNumbers(nums []string) bool {
	// bruteforce approach so not the most efficient but it works
	for i := 0; i < len(nums); i++ {
		rest := make([]string, len(nums)-1)
		copy(rest, nums[:i])
		copy(rest[i:], nums[i+1:])

		if checkNumbers(rest) {
			return true
		}
	}

	return false
}

func checkNumbers(nums []string) bool {
	numbers := make([]int, len(nums))

	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Panicf("Invalid input: %s", num)
		}
		numbers[i] = n
	}

	last := 0
	ascending := true
	for i, num1 := range numbers {
		if i == 0 {
			last = num1
			if numbers[i+1] < num1 {
				ascending = false
			}
			continue
		}

		if ascending && num1 < last {
			return false
		}

		if !ascending && num1 > last {
			return false
		}

		if utils.Abs(num1-last) > 3 || utils.Abs(num1-last) == 0 {
			return false
		}

		last = num1
	}

	return true
}
