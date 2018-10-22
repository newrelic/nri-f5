package entities

import (
	"sync"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-f5/src/client"
)

// CollectVirtualServers collects virtual server entities from F5 and adds them to the integration
func CollectVirtualServers(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup) {
	defer wg.Done()
}
