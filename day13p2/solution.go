package day13p2

import (
	"io"
	"regexp"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	numregex := regexp.MustCompile(`\d+`)

	// create blocks - blocks are 3 lines (until \n)
	blocks := make([]string, 0)
	for i := 0; i < len(lines); i += 4 {
		blocks = append(blocks, lines[i]+"\n"+lines[i+1]+"\n"+lines[i+2])
	}

	sum := 0

	for _, block := range blocks {
		resultsStr := numregex.FindAllString(block, -1)
		// convert to int
		results := make([]int, len(resultsStr))
		for i, r := range resultsStr {
			results[i], _ = strconv.Atoi(r)
		}
		ax, ay, bx, by, px, py := results[0], results[1], results[2], results[3], results[4], results[5]
		px += 10000000000000
		py += 10000000000000
		// i got this by making an equation system
		denom := ax*by - ay*bx
		a := (px*by - py*bx) / denom
		b := (ax*py - ay*px) / denom
		if ax*a+bx*b == px && ay*a+by*b == py {
			sum += a*3 + b
		}
	}

	return sum
}
