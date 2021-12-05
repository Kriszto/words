package scrmabledstrings

import (
	"bufio"
	"github.com/rs/zerolog/log"
	"io"
)

type Input struct {
	reader io.Reader
	dict   *Dictionary
	param  string
}

func NewInput(filename string, options ...func(input *Input)) *Input {
	file := openFile(filename)
	obj := &Input{reader: file}
	for _, option := range options {
		option(obj)
	}

	return obj
}

// WithInputParam sets the threadCount parameter of a Dictionary
func WithDictionary(dict *Dictionary) func(f *Input) {
	return func(i *Input) {
		i.dict = dict
	}
}

// WithInputParam sets the threadCount parameter of a Dictionary
func WithInputParam(param string) func(f *Input) {
	return func(p *Input) {
		p.param = param
	}
}

// ProcessInput takes a pointer to the processed ticket slice,
// reads input in a loop, and prints out the result
func (inp *Input) ProcessInput() []int {
	log.Info().Msg("Reading input")
	ret := make([]int, 0)
	scanner := bufio.NewScanner(inp.reader)
	for scanner.Scan() {
		s := scanner.Text()
		n, _ := inp.dict.Result(s)
		ret = append(ret, n)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal().Err(err).Msg("Scanner error")
	}
	return ret
}
