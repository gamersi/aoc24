package day22p1

import (
	"io"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	sum := 0
	for _, line := range lines {
		initialSecretNum, _ := strconv.Atoi(line)
		newSecretNum := initialSecretNum
		for i := 0; i < 2000; i++ {
			newSecretNum = evolveSecretNum(newSecretNum)
		}
		sum += newSecretNum
	}

	return sum
}

func evolveSecretNum(initialSecretNum int) int {
	secretNum := initialSecretNum
	secretNum = prune(mix(secretNum, secretNum*64))
	secretNum = prune(mix(secretNum, secretNum/32))
	secretNum = prune(mix(secretNum, secretNum*2048))
	return secretNum
}

func mix(a, b int) int {
	return a ^ b
}

func prune(a int) int { // modulo 16777216 (2^24)
	return a & 0xFFFFFF
}
