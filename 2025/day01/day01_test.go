package day01

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed sample.txt
	sampleData string
	//go:embed input.txt
	inputData string
)

func TestSolvePart1_Sample(t *testing.T) {
	in := bytes.NewBuffer([]byte(sampleData))
	out := &bytes.Buffer{}

	SolvePart1(in, out)

	require.Equal(t, "3", out.String())
}

func TestSolvePart1_Input(t *testing.T) {
	in := bytes.NewBuffer([]byte(inputData))
	out := &bytes.Buffer{}

	SolvePart1(in, out)

	require.Equal(t, "1074", out.String())
}

func TestSolvePart2_Sample(t *testing.T) {
	in := bytes.NewBuffer([]byte(sampleData))
	out := &bytes.Buffer{}

	SolvePart2(in, out)

	require.Equal(t, "6", out.String())
}

func TestSolvePart2_Input(t *testing.T) {
	in := bytes.NewBuffer([]byte(inputData))
	out := &bytes.Buffer{}

	SolvePart2(in, out)

	require.Equal(t, "6254", out.String())
}
