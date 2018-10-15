package entities

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func Test_matchesFilter(t *testing.T) {
	testCases := []struct {
		name        string
		filter      []*regexp.Regexp
		shouldMatch bool
	}{
		{
			"test",
			[]*regexp.Regexp{
				regexp.MustCompile("test"),
			},
			true,
		},
		{
			"test",
			[]*regexp.Regexp{
				regexp.MustCompile(".*"),
			},
			true,
		},
		{
			"test",
			[]*regexp.Regexp{
				regexp.MustCompile("tasty"),
			},
			false,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.shouldMatch, matchesFilter(tc.name, tc.filter))
	}
}
