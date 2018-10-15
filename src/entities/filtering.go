package entities

import (
	"regexp"
)

func matchesFilter(name string, filter []*regexp.Regexp) bool {
	for _, pattern := range filter {
		if pattern.MatchString(name) {
			return true
		}
	}

	return false
}
