package util

import (
	"bufio"
	"bytes"
	"io"
	"iter"
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

func ScanSeq(in io.Reader, splitFunc bufio.SplitFunc) iter.Seq[string] {
	return func(yield func(string) bool) {
		scanner := bufio.NewScanner(in)
		scanner.Split(splitFunc)
		for scanner.Scan() {
			line := scanner.Text()

			if !yield(line) {
				return
			}
		}
	}
}

func Lines(in io.Reader) iter.Seq[string] {
	return ScanSeq(in, bufio.ScanLines)
}

func MapIter[From, To any](seq iter.Seq[From], fn func(From) To) iter.Seq[To] {
	return func(yield func(To) bool) {
		for item := range seq {
			res := fn(item)

			if !yield(res) {
				return
			}
		}
	}
}

func ParseLines[T any](in io.Reader, parse func(string) T) iter.Seq[T] {
	lines := Lines(in)
	return MapIter(lines, parse)
}

func ParseSeq[T any](in io.Reader, splitFunc bufio.SplitFunc, parse func(string) T) iter.Seq[T] {
	strs := ScanSeq(in, splitFunc)
	return MapIter(strs, parse)
}

func Modulo(n, mod int) int {
	n %= mod
	if n < 0 {
		n += mod
	}
	return n
}

func Sign(n int) int {
	switch {
	case n > 0:
		return 1
	case n < 0:
		return -1
	default:
		return 0
	}
}

func SplitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}

	return 0, nil, nil
}
