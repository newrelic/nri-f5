package entities

import (
	"sync"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/newrelic/nri-f5/src/definition"
)

// CollectVirtualServers collects virtual server entities from F5 and adds them to the integration
func CollectVirtualServers(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup, pathFilter *arguments.PathMatcher) {
	defer wg.Done()

	var ltmVirtual definition.LtmVirtual
	if err := client.Request("/mgmt/tm/ltm/virtual", &ltmVirtual); err != nil {
		log.Error("Failed to collect inventory for virtual server: %s", err.Error())
	}

	var ltmVirtualStats definition.LtmVirtualStats
	if err := client.Request("/mgmt/tm/ltm/virtual/stats", &ltmVirtualStats); err != nil {
		log.Error("Failed to collect metrics for virtual server: %s", err.Error())
	}

	populateVirtualServerInventory(i, ltmVirtual, pathFilter)
	populateVirtualServerMetrics(i, ltmVirtualStats, pathFilter)
}

func populateVirtualServerInventory(i *integration.Integration, ltmVirtual definition.LtmVirtual, pathFilter *arguments.PathMatcher) {
	for _, virtual := range ltmVirtual.Items {
		if !pathFilter.Matches(virtual.FullPath) {
			continue
		}

		virtualEntity, err := i.Entity(virtual.FullPath, "virtualServer")
		if err != nil {
			log.Error("Failed to get entity object for virtual server %s: %s", virtual.Name, err.Error())
		}

		for k, v := range map[string]interface{}{
			"applicationService": virtual.AppService,
			"destination":        virtual.Destination,
			"kind":               virtual.Kind,
			"maxConnections":     virtual.MaxConnections,
			"name":               virtual.Name,
			"pool":               virtual.Pool,
		} {
			err := virtualEntity.SetInventoryItem(k, "value", v)
			if err != nil {
				log.Error("Failed to set inventory item for %s: %s", k, err.Error())
			}

		}
	}
}

func populateVirtualServerMetrics(i *integration.Integration, ltmVirtualStats definition.LtmVirtualStats, pathFilter *arguments.PathMatcher) {
	for _, virtual := range ltmVirtualStats.Entries {

		entries := virtual.NestedStats.Entries
		virtualName := entries.TmName.Description
		if !pathFilter.Matches(virtualName) {
			continue
		}

		virtualEntity, err := i.Entity(virtualName, "virtualServer")
		if err != nil {
			log.Error("Failed to get entity object for virtual server %s: %s", virtualName, err.Error())
		}

		entries.AvailabilityState.ProcessedDescription = convertAvailabilityState(entries.AvailabilityState.Description)
		entries.EnabledState.ProcessedDescription = convertEnabledState(entries.EnabledState.Description)
		dataIn := entries.DataIn.Value / 8
		dataOut := entries.DataIn.Value / 8
		ephemeralBytesIn := entries.EphemeralBytesIn.Value / 8
		ephemeralBytesOut := entries.EphemeralBytesOut.Value / 8
		entries.DataIn.ProcessedValue = &dataIn
		entries.DataOut.ProcessedValue = &dataOut
		entries.EphemeralBytesIn.ProcessedValue = &ephemeralBytesIn
		entries.EphemeralBytesOut.ProcessedValue = &ephemeralBytesOut

		ms := virtualEntity.NewMetricSet("F5BigIpVirtualServerSample",
			metric.Attribute{Key: "displayName", Value: virtualName},
			metric.Attribute{Key: "entityName", Value: "virtualServer:" + virtualName},
		)

		err = ms.MarshalMetrics(entries)
		if err != nil {
			log.Error("Failed to populate metrics for virtual server %s: %s", virtualName, err)
		}
	}
}
