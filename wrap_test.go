// Copyright (c) 2020 by Gilbert Ramirez <gram@alumni.rice.edu>

// vim: set fileencoding=utf-8 :

package unicodemonowidth

import (
	. "gopkg.in/check.v1"
)

func (s *MySuite) TestNWSIteratorSimple(c *C) {
	input := "Simple string of words."

	it := NewNWSIterator(input)

	c.Check(it.Next(), Equals, "Simple")
	c.Check(it.Next(), Equals, "string")
	c.Check(it.Next(), Equals, "of")
	c.Check(it.Next(), Equals, "words.")
	c.Check(it.Next(), Equals, "")
}

func (s *MySuite) TestNWSIteratorTrim(c *C) {
	input := `
	Trim on the left,
	in   between, and at the end.
	`

	it := NewNWSIterator(input)

	c.Check(it.Next(), Equals, "Trim")
	c.Check(it.Next(), Equals, "on")
	c.Check(it.Next(), Equals, "the")
	c.Check(it.Next(), Equals, "left,")
	c.Check(it.Next(), Equals, "in")
	c.Check(it.Next(), Equals, "between,")
	c.Check(it.Next(), Equals, "and")
	c.Check(it.Next(), Equals, "at")
	c.Check(it.Next(), Equals, "the")
	c.Check(it.Next(), Equals, "end.")
	c.Check(it.Next(), Equals, "")
}

func (s *MySuite) TestWrap(c *C) {

	input := "AAA BB CC DDDDD"

	lines := Wrap(input, 6)

	c.Assert(len(lines), Equals, 3)
	c.Check(lines[0], Equals, "AAA BB")
	c.Check(lines[1], Equals, "CC")
	c.Check(lines[2], Equals, "DDDDD")
}

func (s *MySuite) TestEnglish(c *C) {
	input := `The reason this package
exists is to
help create command-line interface programs, and especially
their help messages. But there could be other uses, too.`

	lines := Wrap(input, 40)
	c.Assert(len(lines), Equals, 5)
	c.Check(lines[0], Equals,
		"The reason this package exists is to")
	c.Check(lines[1], Equals,
		"help create command-line interface")
	c.Check(lines[2], Equals,
		"programs, and especially their help")
	c.Check(lines[3], Equals,
		"messages. But there could be other uses,")
	c.Check(lines[4], Equals,
		"too.")
}
