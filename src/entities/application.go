package entities

import (
	"sync"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/newrelic/nri-f5/src/definition"
)

// CollectApplications collects application entities from F5 and adds them to the integration
func CollectApplications(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup, pathFilter *arguments.PathMatcher, hostPort string, args arguments.ArgumentList) {
	defer wg.Done()

	if !args.HasInventory() {
		return
	}

	var appResponse definition.SysApplicationService
	if err := client.Request("/mgmt/tm/sys/application/service", &appResponse); err != nil {
		log.Error("Couldn't get application service listing from API: %v", err)
		return
	}

	for _, applicationItem := range appResponse.Items {
		if !pathFilter.Matches(applicationItem.FullPath) {
			continue
		}

		applicationPathIDAttr := integration.NewIDAttribute("application", applicationItem.FullPath)
		appEntity, err := i.EntityReportedVia(hostPort, hostPort, "f5-application", applicationPathIDAttr)
		if err != nil {
			log.Error("Couldn't create entity for application object: %v", err)
		}

		for k, v := range map[string]interface{}{
			"deviceGroup":      applicationItem.DeviceGroup,
			"kind":             applicationItem.Kind,
			"name":             applicationItem.Name,
			"template":         applicationItem.Template,
			"templateModified": applicationItem.TemplateModified,
			"trafficGroup":     applicationItem.TrafficGroup,
		} {
			if err := appEntity.SetInventoryItem(k, "value", v); err != nil {
				log.Error("Couldn't set inventory item '%s' on application entity '%s': %v", k, applicationItem.Name, err)
			}
		}

		// find pool to use variable
		for _, variable := range applicationItem.Variables {
			if variable.Name != "pool__pool_to_use" {
				continue
			}

			if err := appEntity.SetInventoryItem("poolToUse", "value", variable.Value); err != nil {
				log.Error("Couldn't set inventory item 'poolToUse' on application entity '%s': %v", applicationItem.Name, err)
			}
		}
	}
}
