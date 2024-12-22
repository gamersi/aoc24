package day22p2

import (
	"io"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	seqToProfit := make(map[[4]int]int)
	visited := make(map[[4]int]bool)
	for _, line := range lines {
		var seq [4]int
		clear(visited)
		initialSecretNum, _ := strconv.Atoi(line)
		newSecretNum := initialSecretNum
		for i := 0; i < 2000; i++ {
			prevPrice := newSecretNum % 10
			newSecretNum = evolveSecretNum(newSecretNum)
			price := newSecretNum % 10
			diff := price - prevPrice
			seq[0], seq[1], seq[2], seq[3] = seq[1], seq[2], seq[3], diff
			if i < 3 {
				continue
			}
			if !visited[seq] {
				seqToProfit[seq] += price
				visited[seq] = true
			}
		}
	}

	maxProfit := 0
	for _, profit := range seqToProfit {
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
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
