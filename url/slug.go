package url

// Google Guidelines on URLS
// https://developers.google.com/search/docs/advanced/guidelines/url-structure

import (
	"regexp"
)

const (
	// Google doesn't really say anything about length, so just something sensible
	// Basically lower case letters and numbers and hyphens (which Google prefers to underscores)
	SlugRegex = "^[a-z0-9]+(?:[:-][a-z0-9]+)*$"
)

var (
	validSlug = regexp.MustCompile(SlugRegex)
)

type Slug string

func (slug Slug) IsValid() bool {
	return validSlug.MatchString(string(slug))
}

func (slug Slug) String() string {
	return string(slug)
}
