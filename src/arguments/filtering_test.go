package arguments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_matchesFilter(t *testing.T) {
	testCases := []struct {
		name        string
		filter      PathMatcher
		shouldMatch bool
	}{
		{
			"/Common/test",
			PathMatcher{
				Partitions: []string{"Common"},
			},
			true,
		},
		{
			"/Common",
			PathMatcher{
				Partitions: []string{"Common"},
			},
			true,
		},
		{
			"/Common",
			PathMatcher{
				Partitions: []string{"Common", "Test"},
			},
			true,
		},
		{
			"/Common/Test",
			PathMatcher{
				Partitions: []string{"Test"},
			},
			false,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.shouldMatch, tc.filter.Matches(tc.name))
	}
}
