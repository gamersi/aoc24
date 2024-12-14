package day14p2

import (
	"io"
	"math"

	"aoc/utils"
)

type Bot struct {
	X, Y   int
	VX, VY int
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	isTest := len(lines) == 12
	if isTest {
		return 0 // no test input for this one
	}
	boardWidth := 101
	boardHeight := 103

	bots := make([]Bot, len(lines))

	for i, line := range lines {
		ints := utils.GetInts(line)
		px, py, vx, vy := ints[0], ints[1], ints[2], ints[3]
		bots[i] = Bot{px, py, vx, vy}
	}

	MAX_MOVES := 10000

	minDoubleOccurences := math.MaxInt
	minStep := 0

	for step := 1; step <= MAX_MOVES; step++ {
		for i, bot := range bots {
			bots[i].X += bot.VX
			bots[i].Y += bot.VY
			if bots[i].X < 0 {
				bots[i].X += boardWidth
			} else if bots[i].X >= boardWidth {
				bots[i].X -= boardWidth
			}
			if bots[i].Y < 0 {
				bots[i].Y += boardHeight
			} else if bots[i].Y >= boardHeight {
				bots[i].Y -= boardHeight
			}
		}
		do := countDoubleOccurences(bots)
		if do < minDoubleOccurences {
			minDoubleOccurences = do
			minStep = step
		}
		if minDoubleOccurences <= 1 {
			break
		}
	}

	return minStep
}

func countDoubleOccurences(bots []Bot) int {
	occurences := make(map[utils.Point]int)

	for _, bot := range bots {
		occurences[utils.Point{X: bot.X, Y: bot.Y}]++
	}

	doubleOccurences := 0
	for _, count := range occurences {
		if count > 1 {
			doubleOccurences++
		}
	}

	return doubleOccurences
}

func printBoard(bots []Bot, boardWidth, boardHeight int) {
	board := make([][]rune, boardHeight)
	for i := range board {
		board[i] = make([]rune, boardWidth)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	for _, bot := range bots {
		board[bot.Y][bot.X] = '#'
	}

	for _, row := range board {
		for _, r := range row {
			print(string(r))
		}
		println()
	}
}
