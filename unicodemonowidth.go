// Copyright (c) 2019 by Gilbert Ramirez <gram@alumni.rice.edu>
package unicodemonowidth

import (
	"golang.org/x/text/unicode/norm"
)

// Returns the count of one-character-wide glyphs that this string
// represents.
func MonoWidth(input string) int {

	s := norm.NFD.String(input)

	count := 0
	for i := 0; i < len(s); {
		d := norm.NFC.NextBoundaryInString(s[i:], true)
		count += 1
		i += d
	}

	return count
}
