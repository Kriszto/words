package scrmabledstrings

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
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
	dir := fmt.Sprintf("testdata/generated/%s", set)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	r, _ := os.Open(fmt.Sprintf("testdata/original/%s_input.txt", set))
	k := 0

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		for i := 0; i < int(num); i++ {
			k++
			color.Green("generating %s #%d", set, k)
			_ = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			words := strings.Fields(scanner.Text())
			scanner.Scan()
			params := strings.Fields(scanner.Text())
			fileInput, err := os.Create(fmt.Sprintf("%s/dictionary_%d.txt", dir, k))
			if err != nil {
				panic(err)
			}
			datawriter := bufio.NewWriter(fileInput)
			for _, data := range words {
				_, _ = datawriter.WriteString(data + "\n")
			}
			datawriter.Flush()
			fileInput.Close()

			fileDict, err := os.Create(fmt.Sprintf("%s/input_%d.txt", dir, k))
			if err != nil {
				panic(err)
			}
			datawriterDict := bufio.NewWriter(fileDict)

			n, _ := strconv.ParseInt(params[2], 10, 64)
			a, _ := strconv.ParseInt(params[3], 10, 64)
			b, _ := strconv.ParseInt(params[4], 10, 64)
			c, _ := strconv.ParseInt(params[5], 10, 64)
			d, _ := strconv.ParseInt(params[6], 10, 64)
			_, _ = datawriterDict.WriteString(g.GenerateInput(
				GeneratorInput{
					S1: rune(params[0][0]),
					S2: rune(params[1][0]),
					N:  n,
					A:  a,
					B:  b,
					C:  c,
					D:  d,
				},
			))

			datawriterDict.Flush()
			fileDict.Close()
		}
	}
}
func (g *Generator) Generate(i GeneratorInput) {

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
