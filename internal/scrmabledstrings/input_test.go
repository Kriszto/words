package scrmabledstrings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput_ReadInput(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
		desc     string
	}{
		{
			input:    "hekjhdfkjhsdfkhj",
			expected: 1,
			desc:     "1-row input",
		},
		{
			input:    "hekjhdfkjhsdfkhj\nkhjgsdkfhgsd",
			expected: 2,
			desc:     "2-row input",
		},
		{
			input:    "",
			expected: 0,
			desc:     "empty input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			p := NewInput(WithInputReader(strings.NewReader(tt.input)))
			res := p.ReadInput()

			assert.Equal(t, tt.expected, len(res))
		})
	}
}

func TestInput_ProcessInput(t *testing.T) {
	var tests = []struct {
		dictionary *Dictionary
		input      []string
		expected   []int
		desc       string
	}{
		{
			dictionary: NewDictionary(WithReader(strings.NewReader("abc\ncde"))),
			input:      []string{"abcde", "abcdd"},
			expected:   []int{2, 1},
			desc:       "2-row dictionary, 2-row input",
		},
		{
			dictionary: NewDictionary(WithReader(strings.NewReader("axpaj\napxaj\ndnrbt\npjxdn\nabd"))),
			input:      []string{"aapxjdnrbtvldptfzbbdbbzxtndrvjblnzjfpvhdhhpxjdnrbt"},
			expected:   []int{4},
			desc:       "5-row dictionary, 1-row input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.dictionary.BuildWords()
			p := NewInput(WithDictionary(tt.dictionary))
			res := p.ProcessInput(tt.input)

			assert.Equal(t, tt.expected, res)
		})
	}
}
