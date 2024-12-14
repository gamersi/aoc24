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
