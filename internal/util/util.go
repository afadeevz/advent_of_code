package util

import (
	"bufio"
	"strconv"
	"strings"
)

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func ParseInt(s string) int {
	x, err := strconv.ParseInt(s, 10, 64)
	CheckErr(err)
	return int(x)
}

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func AssertLen[T any](x []T, l int) {
	if len(x) != l {
		panic("assert len failed")
	}
}

func AssertOdd(x int) {
	if x%2 != 1 {
		panic("assert odd failed")
	}
}
