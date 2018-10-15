package arguments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseError(t *testing.T) {
	testCases := []struct {
		argumentList ArgumentList
		expectError  bool
	}{
		{
			ArgumentList{
				Username:   "",
				Password:   "",
				PathFilter: "[]",
			},
			true,
		},
		{
			ArgumentList{
				Username:   "test",
				Password:   "",
				PathFilter: "[]",
			},
			true,
		},
		{
			ArgumentList{
				Username:   "test",
				Password:   "test",
				PathFilter: "[]",
			},
			false,
		},
		{
			ArgumentList{
				Username:   "test",
				Password:   "test",
				PathFilter: `["test2"]`,
			},
			false,
		},
		{
			ArgumentList{
				Username:   "test",
				Password:   "test",
				PathFilter: `["test2("]`,
			},
			true,
		},
		{
			ArgumentList{
				Username:   "test",
				Password:   "test",
				PathFilter: `["test2"`,
			},
			true,
		},
	}

	for _, tc := range testCases {
		_, err := tc.argumentList.Parse()
		if tc.expectError {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
