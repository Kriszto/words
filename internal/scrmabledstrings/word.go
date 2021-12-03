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
	findings []int
	position int
}

func NewWord(str string, options ...func(dictionary *Word)) *Word {
	f := make([]int, len(str))
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
func WithFindings(findings []int) func(f *Word) {
	return func(p *Word) {
		p.findings = findings
	}
}

// WithPosition sets the findings of a Dictionary
func WithPosition(pos int) func(f *Word) {
	return func(p *Word) {
		p.position = pos
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

	if w.freeLetterPos(l) < 0 && w.hitCount() != len(w.str)-1 {
		w.Reset()
	}

	if w.status == InProgress && w.hitCount() == len(w.str)-1 {
		if l == w.str[len(w.str)-1:len(w.str)] {
			w.position++
			w.findings[len(w.str)-1] = w.position
			w.status = Found
			return
		} else {
			w.Reset()
		}
	}

	if w.status == NoHit && l == w.str[0:1] {
		if len(w.str) == 1 {
			w.status = Found
		} else {
			w.status = InProgress
		}
		w.position++
		w.findings[0] = w.position
	} else if n := w.freeLetterPos(l); w.status == InProgress && n > 0 {
		w.position++
		w.findings[n] = w.position
	}
}

func (w *Word) freeLetterPos(l string) int {
	for i := 0; i < len(w.str)-1; i++ {
		if l == w.str[i:i+1] && w.findings[i] == 0 {
			return i
		}
	}

	return -1
}

func (w *Word) hitCount() int {
	n := 0
	for _, i := range w.findings {
		if i != 0 {
			n++
		}
	}

	return n
}

func (w *Word) fullHit() bool {
	for _, i := range w.findings {
		if i == 0 {
			return false
		}
	}

	return true
}

func (w *Word) Recalc() {
	//originalStr := ""
	w.findings = make([]int, len(w.str))
	w.position = 0
	w.status = NoHit
}

func (w *Word) Reset() {
	w.findings = make([]int, len(w.str))
	w.position = 0
	w.status = NoHit
}

func (w *Word) Equals(w2 *Word) bool {
	return w.status == w2.status
}

func (w *Word) GetStr() string {
	return w.str
}

func (w *Word) GetStatus() WordStatus {
	return w.status
}
