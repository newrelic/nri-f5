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
	// TODO use pathMatcher
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

		virtualEntity, err := i.Entity(virtual.FullPath, "virtualServer") // TODO ensure everywhere is using FullPath as node name
		if err != nil {
			log.Error("Failed to get entity object for virtual server %s: %s", virtual.Name, err.Error())
		}

		logOnError("applicationService", virtual.Name, virtualEntity.SetInventoryItem("applicationService", "value", virtual.AppService))
		logOnError("destination", virtual.Name, virtualEntity.SetInventoryItem("destination", "value", virtual.Destination))
		logOnError("kind", virtual.Name, virtualEntity.SetInventoryItem("kind", "value", virtual.Kind))
		logOnError("maxConnections", virtual.Name, virtualEntity.SetInventoryItem("maxConnections", "value", virtual.MaxConnections))
		logOnError("name", virtual.Name, virtualEntity.SetInventoryItem("name", "value", virtual.Name))
		logOnError("pool", virtual.Name, virtualEntity.SetInventoryItem("pool", "value", virtual.Pool))
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
		entries.DataIn.ProcessedValue = &dataIn
		entries.DataOut.ProcessedValue = &dataOut
		// TODO convert bits to bytes

		ms := virtualEntity.NewMetricSet("F5BigIpVirtualServerSample",
			metric.Attribute{Key: "displayName", Value: virtualName},
			metric.Attribute{Key: "entityType", Value: "virtualServer"},
		)

		err = ms.MarshalMetrics(entries)
		if err != nil {
			log.Error("Failed to populate metrics for virtual server %s: %s", virtualName, err)
		}
	}
}
