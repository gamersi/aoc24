package day17p2

import (
	"io"
	"regexp"
	"slices"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	numberRegexp := regexp.MustCompile(`\d+`)

	instructionsStr := numberRegexp.FindAllString(lines[4], -1)
	instructions := make([]uint64, len(instructionsStr))
	for i := 0; i < len(instructions); i++ {
		instructions[i], _ = strconv.ParseUint(instructionsStr[i], 10, 64)
	}

	return solve(instructions)
}

func solve(instructions []uint64) uint64 {
	aVal := uint64(0)
	for i := len(instructions) - 1; i >= 0; i-- {
		aVal <<= 3
		for !slices.Equal(run(instructions, aVal), instructions[i:]) {
			aVal++
		}
	}
	return aVal
}

func run(instructions []uint64, initialA uint64) []uint64 {
	registers := []uint64{initialA, 0, 0}

	out := []uint64{}

	instrPointer := uint64(0)

	for instrPointer < uint64(len(instructions)) {
		instruction := instructions[instrPointer]
		operand := instructions[instrPointer+1]
		switch instruction {
		case 0:
			registers[0] >>= evalComboOp(operand, registers)
		case 1:
			registers[1] ^= operand
		case 2:
			registers[1] = evalComboOp(operand, registers) & 7
		case 3:
			if registers[0] != 0 {
				instrPointer = operand
				continue
			}
		case 4:
			registers[1] ^= registers[2]
		case 5:
			out = append(out, evalComboOp(operand, registers)&7)
		case 6:
			registers[1] = registers[0] >> evalComboOp(operand, registers)
		case 7:
			registers[2] = registers[0] >> evalComboOp(operand, registers)
		}
		instrPointer += 2
	}
	return out
}

func evalComboOp(instr uint64, registers []uint64) uint64 {
	if instr <= 3 {
		return instr
	}
	if instr <= 6 {
		return registers[instr-4]
	}
	panic("Invalid instruction")
}
