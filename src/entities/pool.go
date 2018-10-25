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

// CollectPools collects pool and pool member entities from F5 and adds them to the integration, using the filter as a whitelist
func CollectPools(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup, partitionFilter *arguments.PathMatcher) {
	defer wg.Done()

	var ltmPool definition.LtmPool
	if err := client.Request("/mgmt/tm/ltm/pool", &ltmPool); err != nil {
		log.Error("Failed to collect inventory for pools: %s", err.Error())
	}

	var ltmPoolStats definition.LtmPoolStats
	if err := client.Request("/mgmt/tm/ltm/pool/stats", &ltmPoolStats); err != nil {
		log.Error("Failed to collect metrics for pools: %s", err.Error())
	}

	populatePoolsInventory(i, ltmPool, ltmPoolStats, partitionFilter)
	populatePoolsMetrics(i, ltmPoolStats, partitionFilter)

	for _, pool := range ltmPool.Items {
		wg.Add(1)
		go func(poolName string) {
			defer wg.Done()
			CollectPoolMembers(poolName, i, client)
		}(pool.FullPath)
	}
}

func populatePoolsInventory(i *integration.Integration, ltmPool definition.LtmPool, ltmPoolStats definition.LtmPoolStats, partitionFilter *arguments.PathMatcher) {
	for _, pool := range ltmPool.Items {
		if !partitionFilter.Matches(pool.FullPath) {
			continue
		}

		poolEntity, err := i.Entity(pool.FullPath, "pool")
		if err != nil {
			log.Error("Failed to get entity object for pool %s: %s", pool.Name, err.Error())
		}

		for k, v := range map[string]interface{}{
			"description":     pool.Description,
			"kind":            pool.Kind,
			"currentLoadMode": pool.LoadBalancingMode,
		} {
			err := poolEntity.SetInventoryItem(k, "value", v)
			if err != nil {
				log.Error("Failed to set inventory item %s: %s", k, err.Error())
			}
		}
	}

	for _, pool := range ltmPoolStats.Entries {
		entries := pool.NestedStats.Entries
		poolName := entries.FullPath.Description
		if !partitionFilter.Matches(poolName) {
			continue
		}

		poolEntity, err := i.Entity(poolName, "pool")
		if err != nil {
			log.Error("Failed to get entity object for pool %s: %s", poolName, err.Error())
		}

		for k, v := range map[string]interface{}{
			"maxConnections": entries.MaxConnections.Value,
			"monitorRule":    entries.MonitorRule.Description,
		} {
			err := poolEntity.SetInventoryItem(k, "value", v)
			if err != nil {
				log.Error("Failed to set inventory item %s: %s", k, err.Error())
			}
		}
	}
}

func populatePoolsMetrics(i *integration.Integration, ltmPoolStats definition.LtmPoolStats, partitionFilter *arguments.PathMatcher) {
	for _, pool := range ltmPoolStats.Entries {
		entries := pool.NestedStats.Entries
		poolName := entries.FullPath.Description
		if !partitionFilter.Matches(poolName) {
			continue
		}

		poolEntity, err := i.Entity(poolName, "pool")
		if err != nil {
			log.Error("Failed to get entity object for pool %s: %s", poolName, err.Error())
		}

		entries.AvailabilityState.ProcessedDescription = convertAvailabilityState(entries.AvailabilityState.Description)
		entries.EnabledState.ProcessedDescription = convertEnabledState(entries.EnabledState.Description)
		dataIn := entries.DataIn.Value / 8
		dataOut := entries.DataIn.Value / 8
		entries.DataIn.ProcessedValue = &dataIn
		entries.DataOut.ProcessedValue = &dataOut

		ms := poolEntity.NewMetricSet("F5BigIpPoolSample",
			metric.Attribute{Key: "displayName", Value: poolName},
			metric.Attribute{Key: "entityName", Value: "pool:" + poolName},
		)

		err = ms.MarshalMetrics(entries)
		if err != nil {
			log.Error("Failed to populate metrics for pool %s: %s", poolName, err)
		}
	}
}
