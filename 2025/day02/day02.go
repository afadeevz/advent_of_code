package day02

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/afadeevz/advent_of_code/internal/util"
	"github.com/samber/lo"
)

type Range struct {
	From, To int
}

func SolvePart1(in io.Reader, out io.Writer) {
	ans := 0
	for r := range util.ParseSeq(in, util.SplitComma, parseRange) {
		for id := r.From; id <= r.To; id++ {
			if !isValidIDPart1(id) {
				ans += id
			}
		}
	}

	fmt.Fprint(out, ans)
}

func SolvePart2(in io.Reader, out io.Writer) {
	ans := 0
	for r := range util.ParseSeq(in, util.SplitComma, parseRange) {
		for id := r.From; id <= r.To; id++ {
			if !isValidIDPart2(id) {
				ans += id
			}
		}
	}

	fmt.Fprint(out, ans)
}

func parseRange(data string) Range {
	parts := strings.Split(data, "-")
	return Range{
		From: int(lo.Must(strconv.ParseInt(parts[0], 10, 64))),
		To:   int(lo.Must(strconv.ParseInt(parts[1], 10, 64))),
	}
}

func isValidIDPart1(id int) bool {
	idStr := strconv.FormatInt(int64(id), 10)
	return isValidIDImpl(idStr, 2)
}

func isValidIDPart2(id int) bool {
	idStr := strconv.FormatInt(int64(id), 10)

	for repeats := 2; repeats <= len(idStr); repeats++ {
		if !isValidIDImpl(idStr, repeats) {
			return false
		}
	}

	return true
}

func isValidIDImpl(idStr string, repeats int) bool {
	if len(idStr)%repeats != 0 {
		return true
	}

	seqLen := len(idStr) / repeats

	for i := range seqLen {
		for j := i + seqLen; j < len(idStr); j += seqLen {
			if idStr[i] != idStr[j] {
				return true
			}
		}
	}

	return false
}
