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

// CollectNodes collects node entities from F5 and adds them to the integration
func CollectNodes(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup, pathFilter *arguments.PathMatcher) {
	defer wg.Done()

	var ltmNode definition.LtmNode
	if err := client.Request("/mgmt/tm/ltm/node", &ltmNode); err != nil {
		log.Error("Failed to collect inventory for nodes: %s", err.Error())
	}

	var ltmNodeStats definition.LtmNodeStats
	if err := client.Request("/mgmt/tm/ltm/node/stats", &ltmNodeStats); err != nil {
		log.Error("Failed to collect metrics for nodes: %s", err.Error())
	}

	populateNodesInventory(i, ltmNode, pathFilter)
	populateNodesMetrics(i, ltmNodeStats, pathFilter)
}

func populateNodesInventory(i *integration.Integration, ltmNode definition.LtmNode, pathFilter *arguments.PathMatcher) {
	for _, node := range ltmNode.Items {
		if !pathFilter.Matches(node.FullPath) {
			continue
		}

		nodeEntity, err := i.Entity(node.FullPath, "node") // TODO ensure everywhere is using FullPath as node name
		if err != nil {
			log.Error("Failed to get entity object for node %s: %s", node.Name, err.Error())
		}

		logOnError("fqdn", node.Name, nodeEntity.SetInventoryItem("fqdn", "value", node.FQDN.TMName))
		logOnError("kind", node.Name, nodeEntity.SetInventoryItem("kind", "value", node.Kind))
		logOnError("address", node.Name, nodeEntity.SetInventoryItem("address", "value", node.Address))
		logOnError("maxConnections", node.Name, nodeEntity.SetInventoryItem("maxConnections", "value", node.MaxConnections))
		logOnError("monitorRule", node.Name, nodeEntity.SetInventoryItem("monitorRule", "value", node.MonitorRule))
	}
}

func populateNodesMetrics(i *integration.Integration, ltmNodeStats definition.LtmNodeStats, pathFilter *arguments.PathMatcher) {
	for _, node := range ltmNodeStats.Entries {
		if !pathFilter.Matches(node.NestedStats.Entries.TmName.Description) {
			continue
		}

		entries := node.NestedStats.Entries
		nodeName := entries.TmName.Description
		nodeEntity, err := i.Entity(nodeName, "node")
		if err != nil {
			log.Error("Failed to get entity object for node %s: %s", nodeName, err.Error())
		}

		entries.MonitorStatus.ProcessedDescription = convertMonitorStatus(entries.MonitorStatus.Description)
		entries.EnabledState.ProcessedDescription = convertEnabledState(entries.EnabledState.Description)
		entries.SessionStatus.ProcessedDescription = convertSessionStatus(entries.SessionStatus.Description)
		entries.AvailabilityState.ProcessedDescription = convertAvailabilityState(entries.AvailabilityState.Description)
		dataIn := entries.DataIn.Value / 8
		dataOut := entries.DataIn.Value / 8
		entries.DataIn.ProcessedValue = &dataIn
		entries.DataOut.ProcessedValue = &dataOut

		ms := nodeEntity.NewMetricSet("F5BigIpNodeSample",
			metric.Attribute{Key: "displayName", Value: nodeName},
			metric.Attribute{Key: "entityType", Value: "node"},
		)

		err = ms.MarshalMetrics(entries)
		if err != nil {
			log.Error("Failed to populate metrics for node %s: %s", nodeName, err)
		}
	}
}
