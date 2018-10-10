package arguments

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestParseError(t *testing.T) {
  testCases := []struct{
    argumentList ArgumentList 
    expectError bool
  }{
    {
      ArgumentList{
        Username: "",
        Password: "",
        PoolMemberFilter: "[]",
        NodeFilter: "[]",
      },
      true,
    },
    {
      ArgumentList{
        Username: "test",
        Password: "",
        PoolMemberFilter: "[]",
        NodeFilter: "[]",
      },
      true,
    },
    {
      ArgumentList{
        Username: "test",
        Password: "test",
        PoolMemberFilter: "[]",
        NodeFilter: "[]",
      },
      false,
    },
    {
      ArgumentList{
        Username: "test",
        Password: "test",
        PoolMemberFilter: `["test1"]`,
        NodeFilter: `["test2"]`,
      },
      false,
    },
    {
      ArgumentList{
        Username: "test",
        Password: "test",
        PoolMemberFilter: `["test1"]`,
        NodeFilter: `["test2("]`,
      },
      true,
    },
    {
      ArgumentList{
        Username: "test",
        Password: "test",
        PoolMemberFilter: `["test1)"]`,
        NodeFilter: `["test2"]`,
      },
      true,
    },
    {
      ArgumentList{
        Username: "test",
        Password: "test",
        PoolMemberFilter: `["test1"`,
        NodeFilter: `["test2"]`,
      },
      true,
    },
    {
      ArgumentList{
        Username: "test",
        Password: "test",
        PoolMemberFilter: `["test1"]`,
        NodeFilter: `["test2"`,
      },
      true,
    },
  }
  
  for _, tc := range testCases {
    _, _, err := tc.argumentList.Parse()
    if tc.expectError {
      assert.Error(t, err)
    } else {
      assert.Nil(t, err)
    }
  }

}

