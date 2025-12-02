package day4

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed test4.txt
	testData string
	//go:embed input4.txt
	inputData string
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{{
		name: "test",
		args: testData,
		want: "18",
	}, {
		name: "input",
		args: inputData,
		want: "2401",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte(tt.args))
			got := Part1(buf)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{{
		name: "test",
		args: testData,
		want: "9",
	}, {
		name: "input",
		args: inputData,
		want: "1822",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte(tt.args))
			got := Part2(buf)
			require.Equal(t, tt.want, got)
		})
	}
}
