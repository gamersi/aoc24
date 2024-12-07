package day07p1

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	sum := 0
	for _, line := range lines {
		split := strings.Split(line, ": ")
		res, _ := strconv.Atoi(split[0])
		numsArr := strings.Split(split[1], " ")
		nums := make([]int, len(numsArr))
		for i, num := range numsArr {
			nums[i], _ = strconv.Atoi(num)
		}
		if lineValid(res, nums) {
			sum += res
		}
	}

	return sum
}

const OPERATORS = "+*"

func lineValid(res int, nums []int) bool {
	bruteforceRes := 0
	for i := 0; i < 1<<uint(len(nums)-1); i++ {
		for j := 0; j < 1<<uint(len(nums)-1); j++ {
			bruteforceRes = nums[0]
			for k := 0; k < len(nums)-1; k++ {
				if j&(1<<uint(k)) > 0 {
					bruteforceRes += nums[k+1]
				} else {
					bruteforceRes *= nums[k+1]
				}
			}
			if bruteforceRes == res {
				return true
			}
		}
	}
	return false
}
