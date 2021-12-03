package scrmabledstrings

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBuildWords(t *testing.T) {
	var tests = []struct {
		data      string
		wordCount int
		words     []string
		input     string
		desc      string
	}{
		{
			data:      "apple",
			wordCount: 1,
			words:     []string{"apple"},
			desc:      "1 word, valid input",
		},
		{
			data:      "apple\n",
			wordCount: 1,
			words:     []string{"apple"},
			desc:      "1 word with new line, valid input",
		},
		{
			data:      "apple\nbanana",
			wordCount: 2,
			words:     []string{"apple", "banana"},

			desc: "2 words, valid input",
		},
		{
			data:      "apple\nbanana\norange",
			wordCount: 3,
			words:     []string{"apple", "banana", "orange"},
			desc:      "3 words, valid input",
		},
		{
			data:      "",
			wordCount: 0,
			words:     []string{},
			desc:      "empty input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			r := strings.NewReader(tt.data)
			d := NewDictionary(WithReader(r))
			d.BuildWords()

			assert.Equal(t, tt.wordCount, d.worldCount())
			assert.Equal(t, tt.words, d.worldStrings())
		})
	}
}

func TestAddLettersAndResult(t *testing.T) {
	var tests = []struct {
		words          []*Word
		input          string
		expectedResult int
		wordCount      int
		desc           string
	}{
		{
			words: []*Word{
				NewWord("apple"),
			},
			input:          "apple",
			expectedResult: 1,
			wordCount:      1,
			desc:           "1 word, 1 hit",
		},
		{
			words: []*Word{
				NewWord("apple"),
				NewWord("banana"),
			},
			input:          "appleljhadbanel",
			expectedResult: 1,
			wordCount:      2,
			desc:           "2 word, 1 hit",
		},
		{
			words: []*Word{
				NewWord("apple"),
				NewWord("banana"),
			},
			input:          "appleasdbananakas",
			expectedResult: 2,
			wordCount:      2,
			desc:           "2 word, 2 hit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			d := NewDictionary(WithWords(tt.words))
			d.addLetters(tt.input)
			assert.Equal(t, tt.wordCount, d.worldCount())
			n, _ := d.Result()
			assert.Equal(t, tt.expectedResult, n)
		})
	}
}
