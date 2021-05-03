package str

import (
	"encoding/json"
	"strings"
)

// CleanString is a case-insensitive string with no leading or trailing whitespace.
// The internal representation is lower case
type CleanString string

type CleanStrings []CleanString

func stringToCleanString(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func NewCleanString(s string) CleanString {
	return CleanString(stringToCleanString(s))
}

func (cs CleanString) String() string {
	return stringToCleanString(string(cs))
}

func (cs CleanString) ToUpper() string {
	return strings.ToUpper(cs.String())
}

func (cs CleanString) ToLower() string {
	return strings.ToLower(cs.String())
}

func (cs *CleanString) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*cs = NewCleanString(s)

	return nil
}

func (cs CleanString) MarshalJSON() ([]byte, error) {
	return json.Marshal(cs.String())
}

func (css CleanStrings) StringSlice() (res []string) {
	for _, cs := range css {
		res = append(res, cs.String())
	}

	return
}
