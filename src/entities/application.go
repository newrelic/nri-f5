package entities

import (
	"sync"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-f5/src/client"
)

// CollectApplications collects application entities from F5 and adds them to the integration
func CollectApplications(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup) {
	defer wg.Done()
}
