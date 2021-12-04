package scrmabledstrings

type TextPortion struct {
	w   *Word
	len int
	pos int
}

func NewTextPortion(str string, l int, options ...func(dictionary *TextPortion)) *TextPortion {
	obj := &TextPortion{w: NewWord(str), len: l, pos: -1}
	for _, option := range options {
		option(obj)
	}
	return obj
}

func (tp *TextPortion) Next() bool {
	return tp.pos < len(tp.w.str)-tp.len
}

func (tp *TextPortion) GetNext() *Word {
	tp.pos++
	return NewWord(tp.w.str[tp.pos : tp.pos+tp.len])
}
