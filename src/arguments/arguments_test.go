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
				Username:        "",
				Password:        "",
				PartitionFilter: "[]",
			},
			true,
		},
		{
			ArgumentList{
				Username:        "test",
				Password:        "",
				PartitionFilter: "[]",
			},
			true,
		},
		{
			ArgumentList{
				Username:        "test",
				Password:        "test",
				PartitionFilter: "[]",
			},
			false,
		},
		{
			ArgumentList{
				Username:        "test",
				Password:        "test",
				PartitionFilter: `["test2"]`,
			},
			false,
		},
		{
			ArgumentList{
				Username:        "test",
				Password:        "test",
				PartitionFilter: `["test2]`,
			},
			true,
		},
		{
			ArgumentList{
				Username:        "test",
				Password:        "test",
				PartitionFilter: `["test2"`,
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
