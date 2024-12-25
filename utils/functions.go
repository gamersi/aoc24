// Package utils implements various utility functions and procedures which have
// been useful for previous Advents of Code.
package utils

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

// regex for integers
var intRegex = regexp.MustCompile(`-?\d+`)

// Greatest Common Denominator
func Gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Least Common Multiple
func Lcm(a, b int64) int64 {
	return a * b / Gcd(a, b)
}

// Kernighan's Bit Counting Algorithm
func CountBits(n uint64) int64 {
	var count int64 = 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}

	return count
}

// Check if error is not nil and panic with message if it is.
func Check(e error, format string, a ...any) {
	if e != nil {
		message := fmt.Sprintf(format, a...)
		panic(fmt.Errorf("%s: %s", message, e))
	}
}

// Read all lines from reader. Panic if there is an issue
func ReadLines(r io.Reader) []string {
	result := []string{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	err := scanner.Err()
	Check(err, "error reading lines")

	return result
}

// Read until \n\n and return blocks of lines
func ReadBlocks(r io.Reader) [][]string {
	result := [][]string{}

	scanner := bufio.NewScanner(r)
	block := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			result = append(result, block)
			block = []string{}
		} else {
			block = append(block, line)
		}
	}
	err := scanner.Err()
	Check(err, "error reading blocks")

	if len(block) > 0 {
		result = append(result, block)
	}

	return result
}

func GetInts(s string) []int {
	matches := intRegex.FindAllString(s, -1)
	result := make([]int, len(matches))
	for i, match := range matches {
		result[i] = Atoi(match)
	}

	return result
}

func Atoi(s string) int {
	result, err := strconv.Atoi(s)
	Check(err, "error converting %s to int", s)

	return result
}

func UIntsToStrings(uints []uint64) []string {
	result := make([]string, len(uints))
	for i, val := range uints {
		result[i] = strconv.FormatUint(val, 10)
	}

	return result
}

func IntsToStrings(ints []int) []string {
	result := make([]string, len(ints))
	for i, val := range ints {
		result[i] = strconv.Itoa(val)
	}

	return result
}

func StringsToInts(strings []string) []int {
	result := make([]int, len(strings))
	for i, str := range strings {
		result[i] = Atoi(str)
	}

	return result
}

func StringsToUInts(strings []string) []uint64 {
	result := make([]uint64, len(strings))
	for i, str := range strings {
		result[i], _ = strconv.ParseUint(str, 10, 64)
	}

	return result
}

func RemoveDuplicates[T comparable](slice []T) []T {
	unique := make(map[T]bool)
	result := []T{}

	for _, item := range slice {
		if !unique[item] {
			unique[item] = true
			result = append(result, item)
		}
	}
	return result
}
