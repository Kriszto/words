package scrmabledstrings

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	inputPattern  = "testdata/original/%s_input.txt"
	outputPattern = "testdata/generated/%s"
)

type Generator struct {
}

type GeneratorInput struct {
	S1            rune
	S2            rune
	N, A, B, C, D int64
}

func NewGenerator(options ...func(dictionary *Generator)) *Generator {
	obj := &Generator{}
	for _, option := range options {
		option(obj)
	}

	return obj
}

// ProcessData iterate through the input reader and returns the slice of processed numbers and issues
func (g *Generator) ProcessData(set string) {
	dir := g.checkDir(set)

	r := openFile(fmt.Sprintf(inputPattern, set))
	k := 0

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		for i := 0; i < int(num); i++ {
			k++
			log.Info().Str("set", set).Str("num", strconv.Itoa(k)).Msg("generating test cases")

			words, seeds := g.scanTestCase(scanner)

			g.creataInput(dir, k, words)

			g.createDictionary(dir, k, seeds)
		}
	}
}

func (g *Generator) createDictionary(dir string, k int, seeds []string) {
	fileDict := createFile(fmt.Sprintf("%s/input_%d.txt", dir, k))
	defer fileDict.Close()

	datawriterDict := bufio.NewWriter(fileDict)
	defer datawriterDict.Flush()

	gi := g.convertParams(seeds)
	_, _ = datawriterDict.WriteString(g.GenerateInput(gi))
}

func (g *Generator) creataInput(dir string, k int, words []string) {
	fileInput := createFile(fmt.Sprintf("%s/dictionary_%d.txt", dir, k))
	defer fileInput.Close()

	datawriter := bufio.NewWriter(fileInput)
	defer datawriter.Flush()

	for _, data := range words {
		_, _ = datawriter.WriteString(data + "\n")
	}
}

func (g *Generator) scanTestCase(scanner *bufio.Scanner) (words, params []string) {
	_ = scanner.Text()
	scanner.Scan()
	scanner.Scan()
	words = strings.Fields(scanner.Text())
	scanner.Scan()
	params = strings.Fields(scanner.Text())
	return
}

func (g *Generator) checkDir(set string) string {
	dir := fmt.Sprintf(outputPattern, set)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return dir
}

func (g *Generator) convertParams(params []string) GeneratorInput {
	s1 := rune(params[0][0])
	s2 := rune(params[1][0])
	n, _ := strconv.ParseInt(params[2], 10, 64)
	a, _ := strconv.ParseInt(params[3], 10, 64)
	b, _ := strconv.ParseInt(params[4], 10, 64)
	c, _ := strconv.ParseInt(params[5], 10, 64)
	d, _ := strconv.ParseInt(params[6], 10, 64)
	gi := GeneratorInput{
		S1: s1,
		S2: s2,
		N:  n,
		A:  a,
		B:  b,
		C:  c,
		D:  d,
	}
	return gi
}

func (g *Generator) GenerateInput(i GeneratorInput) string {
	s1 := i.S1
	s2 := i.S2
	ret := string(s1) + string(s2)
	x1 := int64(s1)
	x2 := int64(s2)
	for k := int64(3); k <= i.N; k++ {
		t := i.A*x2 + i.B*x1 + i.C
		xi := t % i.D
		si := rune(97 + (xi % 26))
		ret += string(si)
		x1, x2 = x2, xi
	}
	return ret
}
