package scrmabledstrings

import "github.com/rs/zerolog/log"

type Fragment struct {
	w             *Word
	len           int
	pos           int
	actualPortion *Word
}

func NewFragment(str string, l int, options ...func(dictionary *Fragment)) *Fragment {
	obj := &Fragment{w: NewWord(str), len: l, pos: -1}
	for _, option := range options {
		option(obj)
	}
	return obj
}

func (tp *Fragment) Next() bool {
	return tp.pos < len(tp.w.str)-tp.len
}

func (tp *Fragment) GetNext() *Word {
	if tp.pos+tp.len > len(tp.w.str)-1 {
		return nil
	}

	tp.pos++
	if tp.actualPortion == nil {
		// Generate a new world
		tp.actualPortion = NewWord(tp.w.str[tp.pos:tp.pos+tp.len], WithBuildFrequency())
	} else {
		// Update frequency
		f := tp.actualPortion.frequency

		// Remove the first character
		f[int(tp.w.str[tp.pos-1])-97]--

		// Add the new character
		f[int(tp.w.str[tp.pos+tp.len-1])-97]++

		// New word with calculated frequency
		tp.actualPortion = NewWord(tp.w.str[tp.pos:tp.pos+tp.len], WithFrequency(&f))
	}
	log.Debug().Str("fragment", tp.actualPortion.str).Msg("Getting input fragment")
	return tp.actualPortion
}
