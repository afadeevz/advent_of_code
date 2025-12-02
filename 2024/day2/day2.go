package day2

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/afadeevz/advent_of_code/internal/util"
	"github.com/samber/lo"
)

func Part1(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	var ans int
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, " ")
		strs = lo.Compact(strs)
		vals := lo.Map(strs, func(s string, _ int) int {
			return util.ParseInt(s)
		})

		if isSafe(vals) {
			ans++
		}
	}

	return fmt.Sprint(ans)
}

func isSafe(vals []int) bool {
	if !isIncreasing(vals) && !isDecreasing(vals) {
		return false
	}

	for i := 0; i < len(vals)-1; i++ {
		d := util.Abs(vals[i] - vals[i+1])
		if d < 1 || d > 3 {
			return false
		}
	}

	return true
}

func isIncreasing(vals []int) bool {
	for i := 0; i < len(vals)-1; i++ {
		if vals[i] >= vals[i+1] {
			return false
		}
	}

	return true
}

func isDecreasing(vals []int) bool {
	for i := 0; i < len(vals)-1; i++ {
		if vals[i] <= vals[i+1] {
			return false
		}
	}

	return true
}

func Part2(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	var ans int
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, " ")
		strs = lo.Compact(strs)
		vals := lo.Map(strs, func(s string, _ int) int {
			return util.ParseInt(s)
		})

		if isSafe2(vals) {
			ans++
		}
	}

	return fmt.Sprint(ans)
}

func isSafe2(vals []int) bool {
	if isSafe(vals) {
		return true
	}

	if isSafe2Impl(vals) {
		return true
	}

	slices.Reverse(vals)
	return isSafe2Impl(vals)
}

func isSafe2Impl(vals []int) bool {
	for i := 0; i < len(vals); i++ {
		without := slices.Clone(vals)
		without = slices.Delete(without, i, i+1)
		if isSafe(without) {
			return true
		}
	}

	return false
}
