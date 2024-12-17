package day17p1

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	numberRegexp := regexp.MustCompile(`\d+`)

	aStart, _ := strconv.ParseUint(numberRegexp.FindString(lines[0]), 10, 64)

	instructionsStr := numberRegexp.FindAllString(lines[4], -1)
	instructions := make([]uint64, len(instructionsStr))
	for i := 0; i < len(instructions); i++ {
		instructions[i], _ = strconv.ParseUint(instructionsStr[i], 10, 64)
	}

	out := run(instructions, aStart)

	return strings.Join(utils.UIntsToStrings(out), ",")
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
