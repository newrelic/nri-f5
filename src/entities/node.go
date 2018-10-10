package entities

import (
	"regexp"
	"sync"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-f5/src/client"
)

// CollectNodes collects node entities from F5 and adds them to the integration, using the filter as a whitelist
func CollectNodes(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup, nodeFilter []*regexp.Regexp) {
	defer wg.Done()
}
