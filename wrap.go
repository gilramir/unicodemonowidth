// Copyright (c) 2020 by Gilbert Ramirez <gram@alumni.rice.edu>
package unicodemonowidth

/*
import (
	"fmt"
)
*/

type nwsIterator struct {
	runes []rune
	pos int
	length int
}

// Iterates over a string, returning non-whitespace spans of
// strings.
func NewNWSIterator(text string) *nwsIterator {
	s := &nwsIterator{
		runes: []rune(text),
	}
	s.length = len(s.runes)

	// Find the first non-whitespace
	for s.pos = 0; s.pos < s.length; s.pos++ {
		r := s.runes[s.pos]
		if r == ' ' || r == '\n' || r == '\t' || r == '\r' {
			continue
		} else {
			break
		}
	}

	return s
}

// Returns the next non-whitespace string ("word" in some languages).
// When the input stream is totally consumed, returns "".
func (s *nwsIterator) Next() string {
	start := s.pos

	// End of the input
	if start >= s.length {
		return ""
	}

	var i int
	for i = start; i < s.length; i++ {
		r := s.runes[i]
		if r == ' ' || r == '\n' || r == '\t' || r == '\r' {
			break
		}
	}
	end := i

	// Find the first non-whitespace
	for s.pos = end; s.pos < s.length; s.pos++ {
		r := s.runes[s.pos]
		if r == ' ' || r == '\n' || r == '\t' || r == '\r' {
			continue
		} else {
			break
		}
	}

	return string(s.runes[start:end])
}

type printedWord struct {
	text string
	width int
}

// http://github.com/jaroslov/knuth-plass-thoughts/blob/mbaster/plass.md

type plassEntry struct {
	first int
	last int
	next int
	score int
}

// Split a string into multiple strings, breaking on whitespace,
// and limiting each line to lineWidth. 
// This does not split English words where they could be split by a hyphen,
// nor does it split sentences in languages like Thai which do not
// put spaces between words.

func Wrap(text string, maxWidth int) []string {
	lines := make([]string, 0)
	words := createWords(text)

	line := ""
	llen := 0
	for _, word := range words {
//		fmt.Printf("llen=%d word=%s\n", llen, word.text)
		if llen == 0 {
			line = word.text
			llen = word.width
		} else if llen + 1 + word.width > maxWidth {
			lines = append(lines, line)
			line = word.text
			llen = word.width
		} else {
			line += " " + word.text
			llen += 1 + word.width
		}
	}
	if llen > 0 {
		lines = append(lines, line)
	}

	return lines
}

func createWords(text string) []*printedWord {
	it := NewNWSIterator(text)

	words := make([]*printedWord, 0)
	for item := it.Next(); item != "" ; item = it.Next() {
		words = append(words, &printedWord{
			text: item,
			width: MonoWidth(item),
		})
	}
	return words
}
