package day25p1

import (
	"io"

	"aoc/utils"
)

type Key [5]int
type Lock [5]int

func (k Key) toLock() Lock {
	var lock Lock
	for i, v := range k {
		lock[i] = v
	}
	return lock
}

func Solve(r io.Reader) any {
	blocks := utils.ReadBlocks(r)

	keys := make([]Key, 0)
	locks := make([]Lock, 0)

	for _, block := range blocks {
		isKey := checkKey(block)
		if isKey {
			keys = append(keys, parseKey(block))
		} else {
			locks = append(locks, parseLock(block))
		}
	}

	nonoverlapping := 0

	for _, key := range keys {
		for _, lock := range locks {
			overlapping := false
			for i := 0; i < 5; i++ {
				if key[i]+lock[i] > 5 {
					overlapping = true
					break
				}
			}
			if !overlapping {
				nonoverlapping += 1
			}
		}
	}

	return nonoverlapping
}

func checkKey(block []string) bool {
	if block[0] == "....." {
		return true
	} else if block[0] == "#####" {
		return false
	} else {
		panic("Invalid block")
	}
}

func parseKey(block []string) Key {
	var key Key
	for _, row := range block[1 : len(block)-1] {
		for j, c := range row {
			if c == '#' {
				key[j] += 1
			}
		}
	}
	return key
}

func parseLock(block []string) Lock {
	return parseKey(block).toLock()
}
