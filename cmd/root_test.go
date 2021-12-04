package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	var tests = []struct {
		number int
		count  int
	}{
		{
			number: 1,
			count:  10,
		},
		{
			number: 2,
			count:  4,
		},
		{
			number: 3,
			count:  21,
		},
		{
			number: 4,
			count:  176,
		},
		{
			number: 5,
			count:  340,
		},
		{
			number: 6,
			count:  26,
		},
		{
			number: 7,
			count:  1,
		},
		{
			number: 8,
			count:  589,
		},
		{
			number: 9,
			count:  2,
		},
		{
			number: 10,
			count:  9,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("test case #%d", tt.number), func(t *testing.T) {
			r := process(fmt.Sprintf("../testdata/generated/dictionary_%d.txt", tt.number), fmt.Sprintf("../testdata/generated/input_%d.txt", tt.number))
			assert.Equal(t, []string{fmt.Sprintf("Case #0: %d", tt.count)}, r)
		})
	}

}
