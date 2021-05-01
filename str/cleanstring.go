package str

import "strings"

// CleanString is a case-insensitive string with no leading or trailing whitespace.
// The internal representation is lower case
type CleanString string

func NewCleanString(s string) CleanString {
	return CleanString(strings.ToLower(strings.TrimSpace(s)))
}

func (cs CleanString) String() string {
	return string(cs)
}

func (cs CleanString) ToUpper() string {
	return strings.ToUpper(string(cs))
}

func (cs CleanString) ToLower() string {
	return strings.ToLower(string(cs))
}
