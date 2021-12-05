package scrmabledstrings

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

type Dictionary struct {
	reader io.Reader
	words  []*Word
}

func NewDictionary(options ...func(dictionary *Dictionary)) *Dictionary {
	obj := &Dictionary{}
	for _, option := range options {
		option(obj)
	}

	return obj
}

// WithReader sets the reader of the Dictionary
func WithReader(r io.Reader) func(f *Dictionary) {
	return func(p *Dictionary) {
		p.reader = r
	}
}

// WithFileName sets the reader of the Dictionary
func WithFileName(filename string) func(f *Dictionary) {
	file := openFile(filename)
	return func(p *Dictionary) {
		p.reader = file
	}
}

// WithWords sets the reader of the Dictionary
func WithWords(w []*Word) func(f *Dictionary) {
	return func(p *Dictionary) {
		p.words = w
	}
}

func (d *Dictionary) BuildWords() {
	log.Info().Msg("Building words")
	scanner := bufio.NewScanner(d.reader)
	for scanner.Scan() {
		d.words = append(d.words, NewWord(scanner.Text(), WithBuildFrequency()))
	}
}

func (d *Dictionary) GetWords() []*Word {
	return d.words
}

func (d *Dictionary) Result(s string) (n, l int) {
	log.Info().Str("s", s).Msg("Processing input")
	num := 0
	for i, w := range d.words {
		if w.IsInString(s) {
			log.Debug().Str("word", w.str).Str("position", fmt.Sprintf("%d/%d", i+1, len(d.words))).Msgf("Word found")
			num++
		} else {
			log.Debug().Str("word", w.str).Str("position", fmt.Sprintf("%d/%d", i+1, len(d.words))).Msgf("Word not found")
		}
	}
	return num, len(d.words)
}

func (d *Dictionary) worldCount() int {
	return len(d.words)
}

func (d *Dictionary) worldStrings() []string {
	ret := make([]string, 0)
	for _, w := range d.words {
		ret = append(ret, w.str)
	}
	return ret
}

// opens file
func openFile(filename string) *os.File {
	log.Info().Str("filename", filename).Msg("Opening file")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Unable to open file %s (%s)", filename, err)
		fmt.Println()
		os.Exit(2)
	}

	return file
}
