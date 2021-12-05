package scrmabledstrings

import (
	"github.com/rs/zerolog/log"
)

type Word struct {
	str       string
	frequency [26]int
}

func NewWord(str string, options ...func(dictionary *Word)) *Word {
	obj := &Word{str: str}
	for _, option := range options {
		option(obj)
	}
	return obj
}

// WithFrequency sets the param of the Word
func WithFrequency(f *[26]int) func(w *Word) {
	return func(w *Word) {
		w.frequency = *f
	}
}

// WithBuildFrequency sets the param of the Word
func WithBuildFrequency() func(w *Word) {
	return func(w *Word) {
		w.buildFrequency()
	}
}

// buildFrequency builds up a word's letter frequency
func (w *Word) buildFrequency() {
	for _, s := range w.str {
		if int(s) < 97 {
			log.Debug().Msgf("invalid character %s", string(s))
			continue
		}
		w.frequency[int(s)-97]++
	}
}

// equals calculates if two words are equal based on letter frequency
func (w *Word) equals(w2 *Word) bool {
	return len(w.str) == len(w2.str) && w.str[0] == w2.str[0] && w.str[len(w.str)-1] == w2.str[len(w2.str)-1] && w.frequency == w2.frequency
}

// IsInString checks if a string contains the scrambled form of a word
func (w *Word) IsInString(s string) bool {
	tp := NewFragment(s, len(w.str))
	for tp.Next() {
		g := tp.GetNext()
		if w.equals(g) {
			return true
		}
	}

	return false
}
