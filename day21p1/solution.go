package day21p1

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"aoc/utils"
)

type Keypad map[rune]utils.Point

func (k Keypad) String() string {
	runes := make([]rune, 0, len(k))
	for r := range k {
		runes = append(runes, r)
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

var numKeypad = Keypad{
	'7': {X: 0, Y: 0},
	'8': {X: 1, Y: 0},
	'9': {X: 2, Y: 0},
	'4': {X: 0, Y: 1},
	'5': {X: 1, Y: 1},
	'6': {X: 2, Y: 1},
	'1': {X: 0, Y: 2},
	'2': {X: 1, Y: 2},
	'3': {X: 2, Y: 2},
	'X': {X: 0, Y: 3},
	'0': {X: 1, Y: 3},
	'A': {X: 2, Y: 3},
}

var dirKeypad = Keypad{
	'X': {X: 0, Y: 0},
	'^': {X: 1, Y: 0},
	'A': {X: 2, Y: 0},
	'<': {X: 0, Y: 1},
	'v': {X: 1, Y: 1},
	'>': {X: 2, Y: 1},
}

func writeVerticalFirst(dx, dy int) string {
	sb := &strings.Builder{}
	for dy < 0 {
		sb.WriteRune('^')
		dy++
	}
	for dy > 0 {
		sb.WriteRune('v')
		dy--
	}
	for dx < 0 {
		sb.WriteRune('<')
		dx++
	}
	for dx > 0 {
		sb.WriteRune('>')
		dx--
	}
	sb.WriteRune('A')
	return sb.String()
}

func writeHorizontalFirst(dx, dy int) string {
	sb := &strings.Builder{}
	for dx < 0 {
		sb.WriteRune('<')
		dx++
	}
	for dx > 0 {
		sb.WriteRune('>')
		dx--
	}
	for dy < 0 {
		sb.WriteRune('^')
		dy++
	}
	for dy > 0 {
		sb.WriteRune('v')
		dy--
	}
	sb.WriteRune('A')
	return sb.String()
}

func getMoves(dx, dy int) []string {
	moves := []string{writeVerticalFirst(dx, dy)}
	if dx != 0 && dy != 0 {
		moves = append(moves, writeHorizontalFirst(dx, dy))
	}
	return moves
}

func doesMoveOverEmpty(point utils.Point, move string, keypad Keypad) bool {
	for _, r := range move {
		switch r {
		case '^':
			point.Y--
		case 'v':
			point.Y++
		case '<':
			point.X--
		case '>':
			point.X++
		}
		if keypad['X'] == point {
			return true
		}
	}
	return false
}

type MemoizedKey struct {
	input   string
	keypads int
}

type MemoizedValue struct {
	length int
	ok     bool
}

var memoized map[MemoizedKey]MemoizedValue

func dfs(input string, keypads []Keypad) (int, bool) {
	key := MemoizedKey{input: input, keypads: len(keypads)}
	if value, ok := memoized[key]; ok {
		return value.length, value.ok
	}
	point := keypads[0]['A']
	length := 0
	for _, r := range input {
		newPoint := keypads[0][r]
		dx := newPoint.X - point.X
		dy := newPoint.Y - point.Y
		moves := getMoves(dx, dy)
		shortest := -1
		found := false
		for _, move := range moves {
			if doesMoveOverEmpty(point, move, keypads[0]) {
				continue
			}
			if len(keypads) == 1 {
				shortest = len(move)
				found = true
				break
			}
			if candidate, ok := dfs(move, keypads[1:]); ok {
				if !found || candidate < shortest {
					shortest = candidate
					found = true
				}
			}
		}
		if !found {
			memoized[key] = MemoizedValue{length: 0, ok: false}
			return 0, false
		}
		length += shortest
		point = newPoint
	}
	memoized[key] = MemoizedValue{length: length, ok: true}
	return length, true
}

func getSum(input []string) int {
	memoized = make(map[MemoizedKey]MemoizedValue)
	keypads := []Keypad{numKeypad, dirKeypad, dirKeypad}
	sum := 0
	for _, line := range input {
		length, ok := dfs(line, keypads)
		if !ok {
			fmt.Println("No path found for", line)
			return -1
		}
		num, _ := strconv.Atoi(line[:len(line)-1])
		sum += length * num
	}
	return sum
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	return getSum(lines)
}
