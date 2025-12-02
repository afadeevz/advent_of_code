package day4

import (
	"bufio"
	"fmt"
	"io"
)

type Dir struct {
	di int
	dj int
}

var Dirs = []Dir{
	{0, 1},
	{0, -1},
	{1, 1},
	{1, 0},
	{1, -1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

func Part1(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	var lines [][]rune
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}

	textToFind := []rune("XMAS")

	ans := 0
	for i, line := range lines {
		for j := range line {
			for _, dir := range Dirs {
				ti, tj := i, j
				found := true
				for _, r := range textToFind {
					if ti < 0 || tj < 0 || ti >= len(lines) || tj >= len(line) {
						found = false
						break
					}

					if lines[ti][tj] != r {
						found = false
						break
					}

					ti += dir.di
					tj += dir.dj
				}

				if found {
					ans++
				}
			}
		}
	}

	return fmt.Sprint(ans)
}

func Part2(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	var lines [][]rune
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}

	sum := 'M' + 'S'

	ans := 0
	for i := 1; i < len(lines)-1; i++ {
		line := lines[i]
		for j := 1; j < len(line)-1; j++ {
			if lines[i][j] != 'A' {
				continue
			}

			if lines[i-1][j-1]+lines[i+1][j+1] != sum {
				continue
			}

			if lines[i+1][j-1]+lines[i-1][j+1] != sum {
				continue
			}

			ans++
		}
	}

	return fmt.Sprint(ans)
}
