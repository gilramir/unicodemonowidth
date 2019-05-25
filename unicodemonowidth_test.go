// Copyright (c) 2019 by Gilbert Ramirez <gram@alumni.rice.edu>

// vim: set fileencoding=utf-8 :

package unicodemonowidth

import (
	. "gopkg.in/check.v1"
)

func (s *MySuite) TestASCII(c *C) {
	c.Check(0, Equals, MonoWidth(""))
	c.Check(1, Equals, MonoWidth("A"))
	c.Check(2, Equals, MonoWidth("AB"))
	c.Check(3, Equals, MonoWidth("XYZ"))
}

func (s *MySuite) TestLatin(c *C) {
	c.Check(0, Equals, MonoWidth(""))
	c.Check(1, Equals, MonoWidth("a"))
	// a with accent
	c.Check(1, Equals, MonoWidth("á"))
}

func (s *MySuite) TestKorean(c *C) {
	// 2 korean characters
	c.Check(2, Equals, MonoWidth("삼성"))
}

func (s *MySuite) TestThai(c *C) {
	// 3 character
	c.Check(3, Equals, MonoWidth("ไหม"))

	//  3 code points, 2 monospaced charaters
	// character with combining diacritic, then another character
	c.Check(2, Equals, MonoWidth("ว่า"))
}
