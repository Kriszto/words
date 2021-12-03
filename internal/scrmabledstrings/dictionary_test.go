package scrmabledstrings

import (
	"os"
	"strings"
	"testing"

	"github.com/rs/zerolog"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	os.Exit(m.Run())
}

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
