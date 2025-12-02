package day5

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/afadeevz/advent_of_code/internal/util"
	"github.com/samber/lo"
)

type Rule struct {
	before int
	after  int
}

func Part1(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	var rules []Rule
	for scanner.Scan() && len(scanner.Text()) > 0 {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		util.AssertLen(parts, 2)

		rules = append(rules, Rule{
			before: util.ParseInt(parts[0]),
			after:  util.ParseInt(parts[1]),
		})
	}

	deps := make(map[int][]int)
	for _, r := range rules {
		deps[r.after] = append(deps[r.after], r.before)
	}

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		pagesStr := strings.Split(line, ",")
		util.AssertOdd(len(pagesStr))
		pages := lo.Map(pagesStr, func(pageStr string, _ int) int {
			return util.ParseInt(pageStr)
		})

		prev := make(map[int]struct{})
		pagesSet := make(map[int]struct{})

		for _, page := range pages {
			pagesSet[page] = struct{}{}
		}

		ok := true
		for _, page := range pages {
			for _, dep := range deps[page] {
				if _, found := pagesSet[dep]; !found {
					continue
				}

				if _, found := prev[dep]; !found {
					ok = false
					goto failed
				}
			}
			prev[page] = struct{}{}
		}
	failed:

		if ok {
			ans += pages[len(pages)/2]
		}
	}

	return fmt.Sprint(ans)
}

func Part2(input io.Reader) string {
	return ""
}
