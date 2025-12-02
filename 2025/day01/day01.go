package day01

import (
	"fmt"
	"io"
	"strconv"

	"github.com/afadeevz/advent_of_code/internal/util"
	"github.com/samber/lo"
)

func SolvePart1(in io.Reader, out io.Writer) {
	dial := 50

	ans := 0
	for rotation := range util.ParseLines(in, parseLine) {
		dial += rotation
		dial = util.Modulo(dial, 100)
		if dial == 0 {
			ans++
		}
	}

	fmt.Fprint(out, ans)
}

func SolvePart2(in io.Reader, out io.Writer) {
	dial := 50

	ans := 0
	for rotation := range util.ParseLines(in, parseLine) {
		for rotation != 0 {
			dial += util.Sign(rotation)
			dial = util.Modulo(dial, 100)
			if dial == 0 {
				ans++
			}
			rotation -= util.Sign(rotation)
		}
	}

	fmt.Fprint(out, ans)
}

func parseLine(line string) int {
	dist := int(lo.Must(strconv.ParseInt(line[1:], 10, 64)))
	dist *= parseDirection(line[0])

	return dist
}

func parseDirection(b byte) int {
	switch b {
	case 'L':
		return -1
	case 'R':
		return 1
	default:
		panic("invalid direction")
	}
}
