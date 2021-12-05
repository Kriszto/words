package scrmabledstrings

import (
	"bufio"
	"io"

	"github.com/rs/zerolog/log"
)

type Input struct {
	reader io.Reader
	dict   *Dictionary
}

func NewInput(options ...func(input *Input)) *Input {
	i := &Input{}
	for _, option := range options {
		option(i)
	}
	return i
}

// WithDictionary sets the Dictionary field of an Input
func WithDictionary(dict *Dictionary) func(f *Input) {
	return func(i *Input) {
		i.dict = dict
	}
}

// WithInputReader sets the reader field of an Input
func WithInputReader(r io.Reader) func(f *Input) {
	return func(i *Input) {
		i.reader = r
	}
}

// WithInputFileName sets the reader field of an Input by filename
func WithInputFileName(filename string) func(f *Input) {
	file := openFile(filename)
	return func(i *Input) {
		i.reader = file
	}
}

// ReadInput takes a pointer to the processed ticket slice,
// reads input in a loop, and prints out the result
func (inp *Input) ReadInput() []string {
	log.Info().Msg("Reading input")
	ret := make([]string, 0)
	scanner := bufio.NewScanner(inp.reader)
	for scanner.Scan() {
		s := scanner.Text()
		ret = append(ret, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal().Err(err).Msg("Scanner error")
	}
	return ret
}

// ProcessInput takes a pointer to the processed ticket slice,
// reads input in a loop, and prints out the result
func (inp *Input) ProcessInput(t []string) []int {
	log.Info().Msg("Processing input")
	ret := make([]int, 0)
	for _, s := range t {
		n, _ := inp.dict.Result(s)
		ret = append(ret, n)
	}

	return ret
}
