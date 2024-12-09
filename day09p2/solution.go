package day09p2

import (
	"io"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	line := lines[0]

	return chksum(defrag(line))
}

func defrag(input string) []int {
	var result []int

	id := 0

	for i, c := range input {
		convInt, _ := strconv.Atoi(string(c))
		storageblock := i%2 == 0
		for j := 0; j < convInt; j++ {
			if storageblock {
				result = append(result, id)
			} else {
				result = append(result, -1)
			}
		}
		if storageblock {
			id++
		}
	}

	fileCount := id
	for fileID := fileCount - 1; fileID >= 0; fileID-- {
		fileStart, fileEnd := -1, -1
		for i := 0; i < len(result); i++ {
			if result[i] == fileID {
				if fileStart == -1 {
					fileStart = i
				}
				fileEnd = i
			}
		}
		if fileStart == -1 {
			continue
		}

		fileLength := fileEnd - fileStart + 1

		freeStart, freeLength := -1, 0
		for i := 0; i < len(result); i++ {
			if result[i] == -1 {
				if freeStart == -1 {
					freeStart = i
				}
				freeLength++
				if freeLength >= fileLength {
					break
				}
			} else {
				freeStart = -1
				freeLength = 0
			}
		}

		if freeLength >= fileLength && freeStart != -1 && freeStart < fileStart {
			for i := 0; i < fileLength; i++ {
				result[freeStart+i] = fileID
			}
			for i := fileStart; i <= fileEnd; i++ {
				result[i] = -1
			}
		}
	}

	return result
}

func chksum(input []int) int {
	sum := 0

	for i, id := range input {
		if id != -1 {
			sum += i * id
		}
	}

	return sum
}
