package scrmabledstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLetter(t *testing.T) {
	var tests = []struct {
		data     *Word
		expected *Word
		letters  string
		desc     string
	}{
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(InProgress), WithFindings([]bool{true, false, false, false, false})),
			letters:  "1",
			desc:     "1 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(InProgress), WithFindings([]bool{true, true, false, false, false})),
			letters:  "12",
			desc:     "2 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(InProgress), WithFindings([]bool{true, true, true, false, false})),
			letters:  "123",
			desc:     "3 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(InProgress), WithFindings([]bool{true, true, true, true, false})),
			letters:  "1234",
			desc:     "4 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(Found), WithFindings([]bool{true, true, true, true, true})),
			letters:  "12345",
			desc:     "5 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(InProgress), WithFindings([]bool{true, true, false, false, false})),
			letters:  "112",
			desc:     "6 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345"),
			letters:  "112346",
			desc:     "7 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(Found), WithFindings([]bool{true, true, true, true, true})),
			letters:  "123456",
			desc:     "8 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345"),
			letters:  "1123789",
			desc:     "9 row, valid input",
		},
		{
			data:     NewWord("apple"),
			expected: NewWord("apple", WithStatus(Found), WithFindings([]bool{true, true, true, true, true})),
			letters:  "apapple",
			desc:     "10 row, valid input",
		},
		{
			data:     NewWord("a"),
			expected: NewWord("a", WithStatus(Found), WithFindings([]bool{true})),
			letters:  "apapple",
			desc:     "11 row, valid input",
		},
		{
			data:     NewWord("ooi"),
			expected: NewWord("ooi", WithStatus(Found), WithFindings([]bool{true, true, true})),
			letters:  "ooi",
			desc:     "12 row, valid input",
		},
		{
			data:     NewWord("ooi"),
			expected: NewWord("ooi", WithStatus(Found), WithFindings([]bool{true, true, true})),
			letters:  "oooi",
			desc:     "13 row, valid input",
		},
		{
			data:     NewWord("ooi"),
			expected: NewWord("ooi", WithStatus(Found), WithFindings([]bool{true, true, true})),
			letters:  "ooooi",
			desc:     "14 row, valid input",
		},
		{
			data:     NewWord("ooi"),
			expected: NewWord("ooi", WithStatus(Found), WithFindings([]bool{true, true, true})),
			letters:  "oooooi",
			desc:     "15 row, valid input",
		},
		{
			data:     NewWord("oco"),
			expected: NewWord("oco", WithStatus(Found), WithFindings([]bool{true, true, true})),
			letters:  "ooco",
			desc:     "16 row, valid input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.data.AddLetters(tt.letters)
			assert.Equal(t, tt.expected, tt.data)
		})
	}
}
func TestFreeLetterPos(t *testing.T) {
	var tests = []struct {
		data        *Word
		expectedPos int
		letters     string
		nextLetter  string
		desc        string
	}{
		{
			data:        NewWord("12345"),
			expectedPos: 0,
			letters:     "",
			nextLetter:  "1",
			desc:        "1 row, valid input",
		},
		{
			data:        NewWord("12345"),
			expectedPos: 1,
			letters:     "1",
			nextLetter:  "2",
			desc:        "2 row, valid input",
		},
		{
			data:        NewWord("12345"),
			expectedPos: 2,
			letters:     "12",
			nextLetter:  "3",
			desc:        "3 row, valid input",
		},
		{
			data:        NewWord("12345"),
			expectedPos: 3,
			letters:     "123",
			nextLetter:  "4",
			desc:        "4 row, valid input",
		},
		{
			data:        NewWord("12345"),
			expectedPos: 4,
			letters:     "1234",
			nextLetter:  "5",
			desc:        "5 row, valid input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			for _, l := range tt.letters {
				tt.data.AddLetter(string(l))
			}
			assert.Equal(t, tt.expectedPos, tt.data.freeLetterPos(tt.nextLetter))
		})
	}
}

func TestReset(t *testing.T) {
	t.Run("Reset", func(t *testing.T) {
		w := NewWord("apple")
		w.AddLetter("a")
		w.AddLetter("p")
		w.AddLetter("p")
		w.Reset()
		assert.Equal(t, NewWord("apple"), w)
	})
}

func TestHitCount(t *testing.T) {
	t.Run("HitCount", func(t *testing.T) {
		w := NewWord("apple")
		w.AddLetters("app")
		assert.Equal(t, 3, w.hitCount())
		w.Reset()
		w.AddLetters("apple")
		assert.Equal(t, 5, w.hitCount())
		w.Reset()
		w.AddLetters("applx")
		assert.Equal(t, 0, w.hitCount())
	})
}
