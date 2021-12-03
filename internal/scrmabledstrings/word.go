package scrmabledstrings

type WordStatus int

const (
	NoHit WordStatus = iota
	InProgress
	Found
)

type Word struct {
	str      string
	status   WordStatus
	findings []bool
}

func NewWord(str string, options ...func(dictionary *Word)) *Word {
	f := make([]bool, len(str))
	obj := &Word{str: str, findings: f}
	for _, option := range options {
		option(obj)
	}

	return obj
}

// WithStatus sets the status of a Dictionary
func WithStatus(s WordStatus) func(f *Word) {
	return func(p *Word) {
		p.status = s
	}
}

// WithFindings sets the findings of a Dictionary
func WithFindings(findings []bool) func(f *Word) {
	return func(p *Word) {
		p.findings = findings
	}
}

func (w *Word) AddLetters(l string) {
	for _, s := range l {
		w.AddLetter(string(s))
	}
}

func (w *Word) AddLetter(l string) {
	if w.status == Found {
		return
	}

	if w.freeLetterPos(l) < 0 {
		w.Reset()
	}
	if w.status == InProgress && w.hitCount() == len(w.str)-1 {
		if l == w.str[len(w.str)-1:len(w.str)] {
			w.findings[len(w.str)-1] = true
			w.status = Found
			return
		} else {
			w.Reset()
		}
	}

	if w.status == NoHit && l == w.str[0:1] {
		w.status = InProgress
		w.findings[0] = true
		return
	}

	if n := w.freeLetterPos(l); w.status == InProgress && n > 0 {
		w.findings[n] = true
		return
	}
}

func (w *Word) freeLetterPos(l string) int {
	for i := 0; i < len(w.str); i++ {
		if l == w.str[i:i+1] && !w.findings[i] {
			return i
		}
	}

	return -1
}

func (w *Word) hitCount() int {
	n := 0
	for _, b := range w.findings {
		if b {
			n++
		}
	}

	return n
}

func (w *Word) Reset() {
	w.findings = make([]bool, len(w.str))
	w.status = NoHit
}
