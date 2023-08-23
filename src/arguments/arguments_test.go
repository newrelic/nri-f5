package arguments_test

import (
	"testing"

	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/stretchr/testify/assert"
)

func TestParseError(t *testing.T) {
	testCases := []struct {
		name         string
		argumentList arguments.ArgumentList
		expectError  error
	}{
		{
			name:         "user name or pass must be specified",
			argumentList: arguments.ArgumentList{},
			expectError:  arguments.ErrMissingUserOrPass,
		},
		{
			name: "max concurrent request must be positive",
			argumentList: arguments.ArgumentList{
				Username:              "test",
				Password:              "test",
				MaxConcurrentRequests: -1,
			},
			expectError: arguments.ErrNegativeMaxConcurrentRequests,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := tc.argumentList.Parse()
			assert.ErrorIs(t, err, tc.expectError)
		})
	}
}
