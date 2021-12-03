package scrmabledstrings

import (
	"bufio"
	"github.com/fatih/color"
	"io"
	"log"
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

// opens file
func (inp *Input) ProcessFile() {
	color.Green("processing file ...")
	r := bufio.NewReader(inp.reader)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			inp.dict.addLetter(string(c))
		}
	}
}
