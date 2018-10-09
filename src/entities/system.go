package entities

import (
	"sync"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-f5/src/client"
)

// CollectSystem collects the system entity from F5 and adds it to the integration
func CollectSystem(integration *integration.Integration, client *client.F5Client, wg *sync.WaitGroup) {
	defer wg.Done()
}
