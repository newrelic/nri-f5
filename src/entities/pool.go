package entities

import (
	"regexp"
	"sync"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-f5/src/client"
)

// CollectPools collects pool and pool member entities from F5 and adds them to the integration, using the filter as a whitelist
func CollectPools(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup, poolMemberFilter []*regexp.Regexp) {
	defer wg.Done()
}
