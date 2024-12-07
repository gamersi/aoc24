package day07p2

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	sum := 0
	for i, line := range lines {
		lineNormalised := strings.Join(strings.Split(line, ": "), " ")
		numsStr := strings.Split(lineNormalised, " ")
		nums := make([]int, len(numsStr))
		for i, numStr := range numsStr {
			nums[i], _ = strconv.Atoi(numStr)
		}
		if lineValid(nums, 1, 0) {
			sum += nums[0]
		}
		fmt.Println(i+1, "/", len(lines))
	}

	return sum
}

func lineValid(nums []int, i, acc int) bool {
	concat, _ := strconv.Atoi(strconv.Itoa(acc) + strconv.Itoa(nums[i]))
	if i == len(nums)-1 {
		return nums[0] == acc + nums[i] || nums[0] == acc * nums[i] || nums[0] == concat
	}
	return lineValid(nums, i+1, acc+nums[i]) || lineValid(nums, i+1, acc*nums[i]) || lineValid(nums, i+1, concat)
}
