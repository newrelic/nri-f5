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
func CollectPools(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup, partitionFilter *arguments.PathMatcher, hostPort string, args arguments.ArgumentList) {
	defer wg.Done()

	var ltmPool definition.LtmPool
	if err := client.Request("/mgmt/tm/ltm/pool", &ltmPool); err != nil {
		log.Error("Failed to collect inventory for pools: %s", err.Error())
	}

	var ltmPoolStats definition.LtmPoolStats
	if err := client.Request("/mgmt/tm/ltm/pool/stats", &ltmPoolStats); err != nil {
		log.Error("Failed to collect metrics for pools: %s", err.Error())
	}

	if args.HasInventory() {
		populatePoolsInventory(i, ltmPool, ltmPoolStats, partitionFilter, hostPort)
	}
	if args.HasMetrics() {
		populatePoolsMetrics(i, ltmPoolStats, partitionFilter, hostPort)
	}

	for _, pool := range ltmPool.Items {
		wg.Add(1)
		go func(poolName string) {
			defer wg.Done()
			CollectPoolMembers(poolName, i, client, hostPort, args)
		}(pool.FullPath)
	}
}

func populatePoolsInventory(i *integration.Integration, ltmPool definition.LtmPool, ltmPoolStats definition.LtmPoolStats, partitionFilter *arguments.PathMatcher, hostPort string) {
	for _, pool := range ltmPool.Items {
		if !partitionFilter.Matches(pool.FullPath) {
			continue
		}

		poolIDAttr := integration.NewIDAttribute("pool", pool.FullPath)
		poolEntity, err := i.EntityReportedVia(hostPort, hostPort, "f5-pool", poolIDAttr)
		if err != nil {
			log.Error("Failed to get entity object for pool %s: %s", pool.Name, err.Error())
		}

		for k, v := range map[string]interface{}{
			"description":     pool.Description,
			"kind":            pool.Kind,
			"currentLoadMode": pool.LoadBalancingMode,
		} {
			// No error check needed since key names are pre-defined
			_ = poolEntity.SetInventoryItem(k, "value", v)
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
			// No error check needed since key names are pre-defined
			_ = poolEntity.SetInventoryItem(k, "value", v)
		}
	}
}

func populatePoolsMetrics(i *integration.Integration, ltmPoolStats definition.LtmPoolStats, partitionFilter *arguments.PathMatcher, hostPort string) {
	for _, pool := range ltmPoolStats.Entries {
		entries := pool.NestedStats.Entries
		poolName := entries.FullPath.Description
		if !partitionFilter.Matches(poolName) {
			continue
		}

		poolIDAttr := integration.NewIDAttribute("pool", poolName)
		poolEntity, err := i.EntityReportedVia(hostPort, hostPort, "f5-pool", poolIDAttr)
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
