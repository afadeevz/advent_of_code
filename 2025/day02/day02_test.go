package day02

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

	require.Equal(t, "1227775554", out.String())
}

func TestSolvePart1_Input(t *testing.T) {
	in := bytes.NewBuffer([]byte(inputData))
	out := &bytes.Buffer{}

	SolvePart1(in, out)

	require.Equal(t, "18700015741", out.String())
}

func TestSolvePart2_Sample(t *testing.T) {
	in := bytes.NewBuffer([]byte(sampleData))
	out := &bytes.Buffer{}

	SolvePart2(in, out)

	require.Equal(t, "4174379265", out.String())
}

func TestSolvePart2_Input(t *testing.T) {
	in := bytes.NewBuffer([]byte(inputData))
	out := &bytes.Buffer{}

	SolvePart2(in, out)

	require.Equal(t, "20077272987", out.String())
}
