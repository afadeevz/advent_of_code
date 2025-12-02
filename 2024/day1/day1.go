package day1

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

	var list1, list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, " ")
		vals = lo.Compact(vals)

		list1 = append(list1, util.ParseInt(vals[0]))
		list2 = append(list2, util.ParseInt(vals[1]))
	}

	slices.Sort(list1)
	slices.Sort(list2)

	var ans int
	for idx := range list1 {
		ans += util.Abs(list1[idx] - list2[idx])
	}

	return fmt.Sprint(ans)
}

func Part2(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	var list1, list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, " ")
		vals = lo.Compact(vals)

		list1 = append(list1, util.ParseInt(vals[0]))
		list2 = append(list2, util.ParseInt(vals[1]))
	}

	counts := lo.CountValues(list2)

	var ans int
	for idx := range list1 {
		val := list1[idx]
		ans += val * counts[val]
	}

	return fmt.Sprint(ans)
}
