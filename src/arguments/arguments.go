package arguments

import (
	"encoding/json"
	"errors"
	"strings"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
)

// ArgumentList contains all the arguments available for the F5 integration
type ArgumentList struct {
	sdkArgs.DefaultArgumentList
	Hostname              string `default:"localhost" help:"The hostname or IP of the F5 BIG IP device to monitor."`
	Port                  int    `default:"443" help:"The port of the iControl API to connect to."`
	Username              string `default:"" help:"The username to connect to the F5 API with."`
	Password              string `default:"" help:"The password to connect to the F5 API with."`
	Timeout               int    `default:"30" help:"The number of seconds to wait before a request times out."`
	CABundleFile          string `default:"" help:"Alternative Certificate Authority bundle file"`
	CABundleDir           string `default:"" help:"Alternative Certificate Authority bundle directory"`
	PartitionFilter       string `default:"[\"Common\"]" help:"JSON array of partitions to collect"`
	MaxConcurrentRequests int    `default:"10" help:"Maximum number of requests running concurrently"`
	ShowVersion           bool   `default:"false" help:"Print build information and exit"`
	SslInsecureSkipVerify bool   `default:"false" help:"Skip verification of the certificate sent by the host."`
}

// Parse validates and parses out regex patterns from the input arguments
func (al *ArgumentList) Parse() (*PathMatcher, error) {
	if al.Username == "" || al.Password == "" {
		return nil, errors.New("both username and password must be provided")
	}

	if al.MaxConcurrentRequests <= 0 {
		return nil, errors.New("max_concurrent_requests must be a positive integer")
	}

	var partitions []string
	err := json.Unmarshal([]byte(al.PartitionFilter), &partitions)
	if err != nil {
		return nil, err
	}

	if al.CABundleFile == "" && al.CABundleDir == "" {
		return nil, errors.New("CABundleFile or CABundleDir must be specified")
	}

	return &PathMatcher{partitions}, nil

}

// PathMatcher is a struct that determines whether a given path should be collected
type PathMatcher struct {
	Partitions []string
}

// Matches returns true if the given path name should be collected
func (p *PathMatcher) Matches(name string) bool {
	for _, pattern := range p.Partitions {
		if strings.HasPrefix(name, "/"+pattern) {
			return true
		}
	}

	return false
}
