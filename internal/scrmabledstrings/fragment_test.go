package scrmabledstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestText_Next(t *testing.T) {
	t.Run("Next() test #1", func(t *testing.T) {
		text := NewFragment("abcde", 3)
		assert.Equal(t, true, text.Next())
		text.GetNext()
		assert.Equal(t, true, text.Next())
		text.GetNext()
		assert.Equal(t, true, text.Next())
		text.GetNext()
		assert.Equal(t, false, text.Next())
	})

	t.Run("Next() test #2", func(t *testing.T) {
		text := NewFragment("ab", 3)

		assert.Equal(t, false, text.Next())
	})

	t.Run("Next() test #3", func(t *testing.T) {
		text := NewFragment("", 3)

		assert.Equal(t, false, text.Next())
	})
}

func TestText_GetNext(t *testing.T) {
	t.Run("Next() test #1", func(t *testing.T) {
		text := NewFragment("abcde", 3)
		w := text.GetNext()
		assert.NotNil(t, w)
		assert.Equal(t, "abc", w.str)
		w = text.GetNext()
		assert.NotNil(t, w)
		assert.Equal(t, "bcd", w.str)
		w = text.GetNext()
		assert.NotNil(t, w)
		assert.Equal(t, "cde", w.str)
		assert.Nil(t, text.GetNext())
	})

	t.Run("Next() test #2", func(t *testing.T) {
		text := NewFragment("ab", 3)
		assert.Nil(t, text.GetNext())
	})

	t.Run("Next() test #3", func(t *testing.T) {
		text := NewFragment("", 3)
		assert.Nil(t, text.GetNext())
	})
}
