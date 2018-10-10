package entities

import (
	"regexp"
	"sync"
  "regexp"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/newrelic/nri-f5/src/definition"
)

// CollectPools collects pool and pool member entities from F5 and adds them to the integration, using the filter as a whitelist
func CollectPools(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup, poolMemberFilter []*regexp.Regexp) {
	defer wg.Done()

  var ltmPool definition.LtmPool
  if err := client.Request("/mgmt/tm/ltm/pool", &ltmPool); err != nil {
    println(err.Error())
  }

  var ltmPoolStats definition.LtmPoolStats
  if err := client.Request("/mgmt/tm/ltm/pool/stats", &ltmPoolStats); err != nil {
    println(err.Error())
  }

  //populatePoolsInventory(i, ltmPool, ltmPoolStats, poolMemberFilter)
  populatePoolsMetrics(i, ltmPoolStats, poolMemberFilter)
}

func populatePoolsInventory(i *integration.Integration, ltmPool definition.LtmPool, ltmPoolStats definition.LtmPoolStats, poolMemberFilter []*regexp.Regexp) {
  for _, pool := range ltmPool.Items {
    poolEntity, err := i.Entity(pool.FullPath, "pool") // TODO ensure everywhere is using FullPath as pool name
    if err != nil {
      log.Error("Failed to get entity object for pool %s: %s", pool.Name, err.Error())
    }

    // TODO handle errors
    err = poolEntity.SetInventoryItem("Description", "value", pool.Description)
    if err != nil {
      log.Error("Failed to set inventory item: %s", err.Error())
    }
    err = poolEntity.SetInventoryItem("Kind", "value", pool.Kind)
    err = poolEntity.SetInventoryItem("Current Load Mode", "value", pool.LoadBalancingMode)
  }


  for _, pool := range ltmPoolStats.Entries {
    entries := pool.NestedStats.Entries   
    poolName := entries.Name.Description
    poolEntity, err := i.Entity(poolName, "pool")
    if err != nil {
      log.Error("Failed to get entity object for pool %s: %s", poolName, err.Error())
    }

    // TODO handle errors
    err = poolEntity.SetInventoryItem("Maximum Connections", "value", entries.MaxConnections.Value)
    err = poolEntity.SetInventoryItem("Monitor Rule", "value", entries.MonitorRule.Description)
  }
}


func populatePoolsMetrics(i *integration.Integration, ltmPoolStats definition.LtmPoolStats, poolMemberFilter []*regexp.Regexp) {
  for _, pool := range ltmPoolStats.Entries {
    entries := pool.NestedStats.Entries   
    poolName := entries.Name.Description
    poolEntity, err := i.Entity(poolName, "pool")
    if err != nil {
      log.Error("Failed to get entity object for pool %s: %s", poolName, err.Error())
    }

    ms := poolEntity.NewMetricSet("F5BigIpPoolSample",
      metric.Attribute{Key:"displayName", Value: poolName},
      metric.Attribute{Key:"entityType", Value: "pool"},
    )

    err = ms.MarshalMetrics(entries)
    if err != nil {
      log.Error("Failed to populate metrics for pool %s: %s", poolName, err)
    }
  }
}
