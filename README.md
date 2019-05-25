When creating console applications, it's important to know how many "wide"
a string is. That is, assuming a monospaced font, how many
visual "characters" will it use. This is complicated by combining diacritics
in Unicode, which are code points that are displayed on top or below another
character.

This function uses the example code for the NextBoundary function at
https://godoc.org/golang.org/x/text/unicode/norm
