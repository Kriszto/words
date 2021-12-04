package scrmabledstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWord_Equals(t *testing.T) {
	var tests = []struct {
		data1    *Word
		data2    *Word
		expected bool
		desc     string
	}{
		{
			data1:    NewWord("12345"),
			data2:    NewWord("13425"),
			expected: true,
			desc:     "q. equal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.data1.Equals(tt.data2))
		})
	}
}

func TestWord_IsInString(t *testing.T) {
	var tests = []struct {
		w        *Word
		s        string
		expected bool
		desc     string
	}{
		{
			w:        NewWord("abcde"),
			s:        "abcdelkh",
			expected: true,
			desc:     "start equal",
		},
		{
			w:        NewWord("abcde"),
			s:        "dfabcdelh",
			expected: true,
			desc:     "middle equal",
		},
		{
			w:        NewWord("abcde"),
			s:        "wsfdabcde",
			expected: true,
			desc:     "end equal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.w.IsInString(tt.s))
		})
	}
}

/*
func TestAddLetter(t *testing.T) {
	var tests = []struct {
		data     *Word
		expected *Word
		letters  string
		desc     string
	}{
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(InProgress), WithFindings([]int{1, 0, 0, 0, 0}), WithPosition(1)),
			letters:  "1",
			desc:     "1 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(InProgress), WithFindings([]int{1, 2, 0, 0, 0}), WithPosition(2)),
			letters:  "12",
			desc:     "2 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(InProgress), WithFindings([]int{1, 2, 3, 0, 0}), WithPosition(3)),
			letters:  "123",
			desc:     "3 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(InProgress), WithFindings([]int{1, 2, 3, 4, 0}), WithPosition(4)),
			letters:  "1234",
			desc:     "4 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(Found), WithFindings([]int{1, 2, 3, 4, 5}), WithPosition(5)),
			letters:  "12345",
			desc:     "5 row, valid input",
		},
		{
			data:     NewWord("12345"),
			expected: NewWord("12345", WithStatus(InProgress), WithFindings([]int{1, 2, 0, 0, 0}), WithPosition(2)),
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
			expected: NewWord("12345", WithStatus(Found), WithFindings([]int{1, 2, 3, 4, 5}), WithPosition(5)),
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
			expected: NewWord("apple", WithStatus(Found), WithFindings([]int{1, 2, 3, 4, 5}), WithPosition(5)),
			letters:  "apapple",
			desc:     "10 row, valid input",
		},
		{
			data:     NewWord("a"),
			expected: NewWord("a", WithStatus(Found), WithFindings([]int{1}), WithPosition(1)),
			letters:  "apapple",
			desc:     "11 row, valid input",
		},
		{
			data:     NewWord("ooi"),
			expected: NewWord("ooi", WithStatus(Found), WithFindings([]int{1, 2, 3}), WithPosition(3)),
			letters:  "ooi",
			desc:     "12 row, valid input",
		},
		{
			data:     NewWord("ooi"),
			expected: NewWord("ooi", WithStatus(Found), WithFindings([]int{1, 2, 3}), WithPosition(3)),
			letters:  "oooi",
			desc:     "13 row, valid input",
		},
		{
			data:     NewWord("ooi"),
			expected: NewWord("ooi", WithStatus(Found), WithFindings([]int{1, 2, 3}), WithPosition(3)),
			letters:  "ooooi",
			desc:     "14 row, valid input",
		},
		{
			data:     NewWord("ooi"),
			expected: NewWord("ooi", WithStatus(Found), WithFindings([]int{1, 2, 3}), WithPosition(3)),
			letters:  "oooooi",
			desc:     "15 row, valid input",
		},
		{
			data:     NewWord("ooio"),
			expected: NewWord("ooio", WithStatus(InProgress), WithFindings([]int{1, 2, 3, 0}), WithPosition(3)),
			letters:  "oooooi",
			desc:     "16 row, valid input",
		},
		{
			data:     NewWord("ooio"),
			expected: NewWord("ooio", WithStatus(Found), WithFindings([]int{1, 2, 3, 4}), WithPosition(4)),
			letters:  "oooooio",
			desc:     "17 row, valid input",
		},
		{
			data:     NewWord("oco"),
			expected: NewWord("oco", WithStatus(Found), WithFindings([]int{1, 2, 3}), WithPosition(3)),
			letters:  "ooco",
			desc:     "18 row, valid input",
		},
		{
			data:     NewWord("kki"),
			expected: NewWord("kki", WithStatus(InProgress), WithFindings([]int{1, 0, 0}), WithPosition(1)),
			letters:  "kik",
			desc:     "19 row, valid input",
		},
		{
			data:     NewWord("kki"),
			expected: NewWord("kki", WithStatus(NoHit), WithFindings([]int{0, 0, 0}), WithPosition(0)),
			letters:  "uksiykikmkmw",
			desc:     "20 row, valid input",
		},
		{
			data:     NewWord("kki"),
			expected: NewWord("kki", WithStatus(NoHit), WithFindings([]int{0, 0, 0}), WithPosition(0)),
			letters:  "kkeki",
			desc:     "21 row, valid input",
		},
		{
			data:     NewWord("bjthd"),
			expected: NewWord("bjthd", WithStatus(NoHit), WithFindings([]int{0, 0, 0}), WithPosition(0)),
			letters:  "zzxvlbhjtdjxzvtbxlnr",
			desc:     "21 row, valid input",
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
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.data.AddLetters(tt.letters)
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
	w := NewWord("apple")
	t.Run("HitCount", func(t *testing.T) {
		w.AddLetters("app")
		assert.Equal(t, 3, w.hitCount())
	})
	t.Run("HitCount", func(t *testing.T) {
		w.Reset()
		w.AddLetters("apple")
		assert.Equal(t, 5, w.hitCount())
	})
	t.Run("HitCount", func(t *testing.T) {
		w.Reset()
		w.AddLetters("applx")
		assert.Equal(t, 0, w.hitCount())
	})
}

func TestRewind(t *testing.T) {
	var tests = []struct {
		data     *Word
		expected *Word
		desc     string
	}{
		{
			data:     NewWord("ooi", WithStatus(InProgress), WithFindings([]int{1, 2, 0}), WithPosition(2)),
			expected: NewWord("ooi", WithStatus(InProgress), WithFindings([]int{1, 0, 0}), WithPosition(1)),
			desc:     "1. rewind test",
		},
		{
			data:     NewWord("ooi", WithStatus(Found), WithFindings([]int{1, 2, 3}), WithPosition(3)),
			expected: NewWord("ooi", WithStatus(Found), WithFindings([]int{1, 2, 3}), WithPosition(3)),
			desc:     "2. rewind test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.data.Recalc()
			assert.Equal(t, tt.expected, tt.data)
		})
	}
}*/
