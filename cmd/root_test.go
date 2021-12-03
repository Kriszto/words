package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	os.Exit(m.Run())
}

func TestProcess(t *testing.T) {
	var tests = []struct {
		number int
		count  int
	}{
		{number: 1, count: 10},
		{number: 2, count: 4},
		{number: 3, count: 21},
		{number: 4, count: 176},
		{number: 5, count: 340},
		{number: 6, count: 26},
		{number: 7, count: 1},
		{number: 8, count: 589},
		{number: 9, count: 2},
		{number: 10, count: 9},
		{number: 11, count: 14},
		{number: 12, count: 68},
		{number: 13, count: 1},
		{number: 14, count: 2},
		{number: 15, count: 126},
		{number: 16, count: 31},
		{number: 17, count: 3},
		{number: 18, count: 2},
		{number: 19, count: 447},
		{number: 20, count: 168},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("test case #%d", tt.number), func(t *testing.T) {
			r := process(
				fmt.Sprintf("../testdata/generated/ts1/dictionary_%d.txt", tt.number),
				fmt.Sprintf("../testdata/generated/ts1/input_%d.txt", tt.number),
			)
			assert.Equal(t, []string{fmt.Sprintf("Case #0: %d", tt.count)}, r)
		})
	}
}
