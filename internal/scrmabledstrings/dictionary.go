package scrmabledstrings

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Ak-Army/xlog"
	"github.com/fatih/color"
)

type Dictionary struct {
	reader io.Reader
	words  []*Word
	param  string
}

func NewDictionary(options ...func(dictionary *Dictionary)) *Dictionary {
	obj := &Dictionary{}
	for _, option := range options {
		option(obj)
	}

	return obj
}

// WithParam sets the param of the Dictionary
func WithParam(param string) func(f *Dictionary) {
	return func(p *Dictionary) {
		p.param = param
	}
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
	color.Green("building....")
	scanner := bufio.NewScanner(d.reader)

	const maxCapacity = 100000000 // your required line length
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		d.words = append(d.words, NewWord(scanner.Text()))
	}
}

func (d *Dictionary) GetWords() []*Word {
	return d.words
}

func (d *Dictionary) worldCount() int {
	return len(d.words)
}

func (d *Dictionary) worldStrings() (ret []string) {
	ret = make([]string, 0)
	for _, w := range d.words {
		ret = append(ret, w.str)
	}
	return
}

func (d *Dictionary) Debug(s string) {
	n, l := d.Result(s)
	xlog.Debugf("result %d/%d", n, l)
}

func (d *Dictionary) Result(s string) (n, l int) {
	num := 0
	for _, w := range d.words {
		if w.IsInString(s) {
			num++
		}
	}
	return num, len(d.words)
}

// opens file
func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Unable to open file %s (%s)", filename, err)
		fmt.Println()
		os.Exit(2)
	}

	return file
}
