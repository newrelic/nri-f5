package arguments

import (
  "regexp"
  "errors"
  "encoding/json"
	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
)

// ArgumentList contains all the arguments available for the F5 integration
type ArgumentList struct {
	sdkArgs.DefaultArgumentList
	Hostname         string `default:"localhost" help:"The hostname or IP of the F5 BIG IP device to monitor."`
	Port             int    `default:"443" help:"The port of the iControl API to connect to."`
	Username         string `default:"" help:"The username to connect to the F5 API with."`
	Password         string `default:"" help:"The password to connect to the F5 API with."`
	Timeout          int    `default:"30" help:"The number of seconds to wait before a request times out."`
	UseSSL           bool   `default:"true" help:"Whether or not to use SSL to connect to the API. The F5 API only allows connections using SSL."`
	CABundleFile     string `default:"" help:"Alternative Certificate Authority bundle file"`
	CABundleDir      string `default:"" help:"Alternative Certificate Authority bundle directory"`
	PoolMemberFilter string `default:"[]" help:"JSON array of pool member name regexes to collect."`
	NodeFilter       string `default:"[]" help:"JSON array of node name regexes to collect."`
}

// Parse validates and parses out regex patterns from the input arguments
func (al *ArgumentList) Parse() ([]*regexp.Regexp, []*regexp.Regexp, error) {
  if al.Username == "" || al.Password == "" {
    return nil, nil, errors.New("both username and password must be provided")
  }

  var poolMemberRegexStrings []string
  if err := json.Unmarshal([]byte(al.PoolMemberFilter), &poolMemberRegexStrings); err != nil {
    return nil, nil, err
  }

  poolMemberRegexPatterns := make([]*regexp.Regexp, len(poolMemberRegexStrings))
  for i, regexString := range poolMemberRegexStrings {
    pattern, err := regexp.Compile(regexString)
    if err != nil {
      return nil, nil, err
    }
    poolMemberRegexPatterns[i] = pattern
  }

  var nodeRegexStrings []string
  if err := json.Unmarshal([]byte(al.NodeFilter), &nodeRegexStrings); err != nil {
    return nil, nil, err
  }

  nodeRegexPatterns := make([]*regexp.Regexp, len(nodeRegexStrings))
  for i, regexString := range nodeRegexStrings {
    pattern, err := regexp.Compile(regexString)
    if err != nil {
      return nil, nil, err
    }
    nodeRegexPatterns[i] = pattern
  }

  return poolMemberRegexPatterns, nodeRegexPatterns, nil

}
